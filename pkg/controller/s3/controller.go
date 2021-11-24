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

package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/fhopfensperger/provider-basic-aws/apis/s3/v1beta1"
	awsclient "github.com/fhopfensperger/provider-basic-aws/pkg/clients"
	"github.com/fhopfensperger/provider-basic-aws/pkg/clients/s3"
)

// SetupBucket adds a controller that reconciles Buckets.
func SetupBucket(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(v1beta1.BucketGroupKind)
	logger := l.WithValues("controller", name)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewController(rl),
		}).
		For(&v1beta1.Bucket{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1beta1.BucketGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), newClientFn: s3.NewClient, logger: logger}),
			managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
			managed.WithPollInterval(poll),
			managed.WithLogger(logger),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type external struct {
	kube     client.Client
	s3client s3.BucketClient
	logger   logging.Logger
}

const (
	errUnexpectedObject = "The managed resource is not a Bucket"
	errHead             = "failed to query Bucket"
	errCreate           = "failed to create the Bucket"
)

type connector struct {
	kube        client.Client
	newClientFn func(config aws.Config) s3.BucketClient
	logger      logging.Logger
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1beta1.Bucket)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	cfg, err := awsclient.GetConfig(ctx, c.kube, mg, cr.Spec.ForProvider.LocationConstraint)
	if err != nil {
		return nil, err
	}
	s3client := c.newClientFn(*cfg)
	return &external{s3client: s3client, kube: c.kube, logger: c.logger}, nil
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint: gocyclo
	cr, ok := mg.(*v1beta1.Bucket)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}

	_, err := e.s3client.HeadBucket(ctx, &awss3.HeadBucketInput{Bucket: aws.String(meta.GetExternalName(cr))})
	if err != nil {
		return managed.ExternalObservation{}, awsclient.Wrap(resource.Ignore(IsNotFound, err), errHead)
	}

	resp1, err := e.s3client.ListBuckets(ctx, nil)
	if err != nil {
		return managed.ExternalObservation{}, awsclient.Wrap(resource.Ignore(IsNotFound, err), errHead)
	}
	cr.Status.AtProvider = GenerateBucketObservation(meta.GetExternalName(cr), resp1.Buckets)

	cr.Status.SetConditions(xpv1.Available())
	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
		ConnectionDetails: map[string][]byte{
			xpv1.ResourceCredentialsSecretEndpointKey:  []byte(meta.GetExternalName(cr)),
			v1beta1.ResourceCredentialsSecretRegionKey: []byte(cr.Spec.ForProvider.LocationConstraint),
		},
	}, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1beta1.Bucket)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}

	_, err := e.s3client.CreateBucket(ctx, GenerateCreateBucketInput(meta.GetExternalName(cr), cr.Spec.ForProvider))
	if resource.Ignore(IsAlreadyExists, err) != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	return managed.ExternalCreation{}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	_, ok := mg.(*v1beta1.Bucket)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}

	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1beta1.Bucket)
	if !ok {
		return errors.New(errUnexpectedObject)
	}

	cr.Status.SetConditions(xpv1.Deleting())
	_, err := e.s3client.DeleteBucket(ctx, &awss3.DeleteBucketInput{Bucket: aws.String(meta.GetExternalName(cr))})
	return resource.Ignore(IsNotFound, err)
}
