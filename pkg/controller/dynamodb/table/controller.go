/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package table

import (
	"context"
	"encoding/json"
	"time"

	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	svcsdkapi "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapi "github.com/aws/aws-sdk-go/service/dynamodb"
	svcapitypes "github.com/fhopfensperger/provider-basic-aws/apis/dynamodb/v1beta1"
	aws "github.com/fhopfensperger/provider-basic-aws/pkg/clients"
	awsclient "github.com/fhopfensperger/provider-basic-aws/pkg/clients"
)

// SetupTable adds a controller that reconciles Table.
func SetupTable(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(svcapitypes.TableGroupKind)
	logger := l.WithValues("controller", name)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewController(rl),
		}).
		For(&svcapitypes.Table{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.TableGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), logger: logger}),
			managed.WithInitializers(
				managed.NewNameAsExternalName(mgr.GetClient()),
				managed.NewDefaultProviderConfig(mgr.GetClient())),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connector struct {
	kube   client.Client
	logger logging.Logger
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Table)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return &external{c.kube, svcapi.New(sess), c.logger}, nil
}

type external struct {
	kube   client.Client
	client svcsdkapi.DynamoDBAPI
	logger logging.Logger
}

const (
	errUnexpectedObject = "managed resource is not an Table resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Table in AWS"
	errUpdate        = "cannot update Table in AWS"
	errDescribe      = "failed to describe Table"
	errDelete        = "failed to delete Table"
)

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Table)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeTableInput(cr)
	input.TableName = aws.String(meta.GetExternalName(cr))
	resp, err := e.client.DescribeTableWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()

	if err := lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateTable(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}

	if err != nil {
		return managed.ExternalObservation{}, err
	}
	switch aws.StringValue(resp.Table.TableStatus) {
	case string(svcapitypes.TableStatus_SDK_CREATING):
		cr.SetConditions(xpv1.Creating())
	case string(svcapitypes.TableStatus_SDK_DELETING):
		cr.SetConditions(xpv1.Deleting())
	case string(svcapitypes.TableStatus_SDK_ACTIVE):
		cr.SetConditions(xpv1.Available())
	case string(svcapitypes.TableStatus_SDK_ARCHIVED), string(svcapitypes.TableStatus_SDK_INACCESSIBLE_ENCRYPTION_CREDENTIALS), string(svcapitypes.TableStatus_SDK_ARCHIVING):
		cr.SetConditions(xpv1.Unavailable())
	}

	// Check tags
	tagResp, err := e.client.ListTagsOfResource(&svcsdk.ListTagsOfResourceInput{ResourceArn: resp.Table.TableArn})
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	if len(tagResp.Tags) > 0 {
		currentSpec.Tags = nil
		for _, tag := range tagResp.Tags {
			currentSpec.Tags = append(currentSpec.Tags, &svcapitypes.Tag{Value: tag.Value, Key: tag.Key})
		}
	}

	return managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate && cmp.Equal(cr.Spec.ForProvider.Tags, currentSpec.Tags),
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
		ConnectionDetails: managed.ConnectionDetails{
			"tableName":         []byte(meta.GetExternalName(cr)),
			"tableArn":          []byte(aws.StringValue(resp.Table.TableArn)),
			"latestStreamArn":   []byte(aws.StringValue(resp.Table.LatestStreamArn)),
			"latestStreamLabel": []byte(aws.StringValue(resp.Table.LatestStreamLabel)),
		},
	}, nil
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Table)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateTableInput(cr)
	input.TableName = aws.String(meta.GetExternalName(cr))
	resp, err := e.client.CreateTableWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.TableDescription.TableArn != nil {
		cr.Status.AtProvider.TableARN = resp.TableDescription.TableArn
	} else {
		cr.Status.AtProvider.TableARN = nil
	}
	if resp.TableDescription.TableId != nil {
		cr.Status.AtProvider.TableID = resp.TableDescription.TableId
	} else {
		cr.Status.AtProvider.TableID = nil
	}
	if resp.TableDescription.TableName != nil {
		cr.Status.AtProvider.TableName = resp.TableDescription.TableName
	} else {
		cr.Status.AtProvider.TableName = nil
	}
	if resp.TableDescription.TableSizeBytes != nil {
		cr.Status.AtProvider.TableSizeBytes = resp.TableDescription.TableSizeBytes
	} else {
		cr.Status.AtProvider.TableSizeBytes = nil
	}
	if resp.TableDescription.TableStatus != nil {
		cr.Status.AtProvider.TableStatus = resp.TableDescription.TableStatus
	} else {
		cr.Status.AtProvider.TableStatus = nil
	}

	return managed.ExternalCreation{
		ConnectionDetails: nil,
	}, nil
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Table)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	var err error
	if len(cr.Spec.ForProvider.Tags) > 0 {
		keys := []*string{}
		tags := []*svcsdk.Tag{}
		for _, tag := range cr.Spec.ForProvider.Tags {
			tags = append(tags, &svcsdk.Tag{Value: tag.Value, Key: tag.Key})
		}

		var tagsResp *svcapi.ListTagsOfResourceOutput
		tagsResp, err = e.client.ListTagsOfResource(&svcapi.ListTagsOfResourceInput{ResourceArn: cr.Status.AtProvider.TableARN})
		if err != nil {
			return managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate)
		}
		for _, existingTag := range tagsResp.Tags {
			keys = append(keys, existingTag.Key)
		}
		_, err = e.client.UntagResource(&svcapi.UntagResourceInput{ResourceArn: cr.Status.AtProvider.TableARN, TagKeys: keys})
		if err != nil {
			return managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate)
		}
		_, err = e.client.TagResource(&svcapi.TagResourceInput{ResourceArn: cr.Status.AtProvider.TableARN, Tags: tags})
	}
	return managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate)
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Table)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteTableInput(cr)
	_, err := e.client.DeleteTableWithContext(ctx, input)
	return awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete)
}

func lateInitialize(in *svcapitypes.TableParameters, t *svcsdk.DescribeTableOutput) error { // nolint:gocyclo,unparam
	if t == nil {
		return nil
	}

	if len(in.AttributeDefinitions) == 0 && len(t.Table.AttributeDefinitions) != 0 {
		in.AttributeDefinitions = buildAttributeDefinitions(t.Table.AttributeDefinitions)
	}

	if in.KeySchema == nil && len(t.Table.KeySchema) != 0 {
		in.KeySchema = buildAlphaKeyElements(t.Table.KeySchema)
	}

	if in.ProvisionedThroughput == nil && t.Table.ProvisionedThroughput != nil {
		in.ProvisionedThroughput = &svcapitypes.ProvisionedThroughput{
			ReadCapacityUnits:  t.Table.ProvisionedThroughput.ReadCapacityUnits,
			WriteCapacityUnits: t.Table.ProvisionedThroughput.WriteCapacityUnits,
		}
	}

	return nil
}

func buildAlphaKeyElements(keys []*svcsdk.KeySchemaElement) []*svcapitypes.KeySchemaElement {
	if len(keys) == 0 {
		return nil
	}
	keyElements := make([]*svcapitypes.KeySchemaElement, len(keys))
	for i, val := range keys {
		keyElements[i] = &svcapitypes.KeySchemaElement{
			AttributeName: val.AttributeName,
			KeyType:       val.KeyType,
		}
	}
	return keyElements
}

func buildAttributeDefinitions(attributes []*svcsdk.AttributeDefinition) []*svcapitypes.AttributeDefinition {
	if len(attributes) == 0 {
		return nil
	}
	attributeDefinitions := make([]*svcapitypes.AttributeDefinition, len(attributes))
	for i, val := range attributes {
		attributeDefinitions[i] = &svcapitypes.AttributeDefinition{
			AttributeName: val.AttributeName,
			AttributeType: val.AttributeType,
		}
	}
	return attributeDefinitions
}

// createPatch creates a *svcapitypes.TableParameters that has only the changed
// values between the target *svcapitypes.TableParameters and the current
// *dynamodb.TableDescription
func createPatch(in *svcsdk.DescribeTableOutput, target *svcapitypes.TableParameters) (*svcapitypes.TableParameters, error) {
	currentParams := &svcapitypes.TableParameters{}
	if err := lateInitialize(currentParams, in); err != nil {
		return nil, err
	}

	jsonPatch, err := aws.CreateJSONPatch(currentParams, target)
	if err != nil {
		return nil, err
	}
	patch := &svcapitypes.TableParameters{}
	if err := json.Unmarshal(jsonPatch, patch); err != nil {
		return nil, err
	}
	return patch, nil
}

func isUpToDate(cr *svcapitypes.Table, resp *svcsdk.DescribeTableOutput) (bool, error) {
	// A table that's currently updating or creating can't be updated, so we
	// temporarily consider it to be up-to-date no matter what.
	switch aws.StringValue(cr.Status.AtProvider.TableStatus) {
	case string(svcapitypes.TableStatus_SDK_UPDATING), string(svcapitypes.TableStatus_SDK_CREATING):
		return true, nil
	}

	_, err := createPatch(resp, &cr.Spec.ForProvider)
	if err != nil {
		return false, err
	}

	return true, nil
}
