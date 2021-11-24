/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	aws "github.com/fhopfensperger/provider-basic-aws/pkg/clients"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/fhopfensperger/provider-basic-aws/apis/dynamodb/v1beta1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeTableInput returns input for read
// operation.
func GenerateDescribeTableInput(cr *svcapitypes.Table) *svcsdk.DescribeTableInput {
	res := &svcsdk.DescribeTableInput{}

	if cr.Status.AtProvider.TableName != nil {
		res.SetTableName(*cr.Status.AtProvider.TableName)
	}

	return res
}

// GenerateTable returns the current state in the form of *svcapitypes.Table.
func GenerateTable(resp *svcsdk.DescribeTableOutput) *svcapitypes.Table {
	cr := &svcapitypes.Table{}

	if resp.Table.CreationDateTime != nil {
		cr.Status.AtProvider.CreationDateTime = &metav1.Time{*resp.Table.CreationDateTime}
	} else {
		cr.Status.AtProvider.CreationDateTime = nil
	}
	if resp.Table.TableArn != nil {
		cr.Status.AtProvider.TableARN = resp.Table.TableArn
	} else {
		cr.Status.AtProvider.TableARN = nil
	}
	if resp.Table.TableId != nil {
		cr.Status.AtProvider.TableID = resp.Table.TableId
	} else {
		cr.Status.AtProvider.TableID = nil
	}
	if resp.Table.TableName != nil {
		cr.Status.AtProvider.TableName = resp.Table.TableName
	} else {
		cr.Status.AtProvider.TableName = nil
	}
	if resp.Table.TableSizeBytes != nil {
		cr.Status.AtProvider.TableSizeBytes = resp.Table.TableSizeBytes
	} else {
		cr.Status.AtProvider.TableSizeBytes = nil
	}
	if resp.Table.TableStatus != nil {
		cr.Status.AtProvider.TableStatus = resp.Table.TableStatus
	} else {
		cr.Status.AtProvider.TableStatus = nil
	}

	return cr
}

// GenerateCreateTableInput returns a create input.
func GenerateCreateTableInput(cr *svcapitypes.Table) *svcsdk.CreateTableInput {
	res := &svcsdk.CreateTableInput{}

	if cr.Spec.ForProvider.AttributeDefinitions != nil {
		f0 := []*svcsdk.AttributeDefinition{}
		for _, f0iter := range cr.Spec.ForProvider.AttributeDefinitions {
			f0elem := &svcsdk.AttributeDefinition{}
			if f0iter.AttributeName != nil {
				f0elem.SetAttributeName(*f0iter.AttributeName)
			}
			if f0iter.AttributeType != nil {
				f0elem.SetAttributeType(*f0iter.AttributeType)
			}
			f0 = append(f0, f0elem)
		}
		res.SetAttributeDefinitions(f0)
	}
	if cr.Spec.ForProvider.KeySchema != nil {
		f3 := []*svcsdk.KeySchemaElement{}
		for _, f3iter := range cr.Spec.ForProvider.KeySchema {
			f3elem := &svcsdk.KeySchemaElement{}
			if f3iter.AttributeName != nil {
				f3elem.SetAttributeName(*f3iter.AttributeName)
			}
			if f3iter.KeyType != nil {
				f3elem.SetKeyType(*f3iter.KeyType)
			}
			f3 = append(f3, f3elem)
		}
		res.SetKeySchema(f3)
	}
	if cr.Spec.ForProvider.ProvisionedThroughput != nil {
		f5 := &svcsdk.ProvisionedThroughput{}
		if cr.Spec.ForProvider.ProvisionedThroughput.ReadCapacityUnits != nil {
			f5.SetReadCapacityUnits(*cr.Spec.ForProvider.ProvisionedThroughput.ReadCapacityUnits)
		}
		if cr.Spec.ForProvider.ProvisionedThroughput.WriteCapacityUnits != nil {
			f5.SetWriteCapacityUnits(*cr.Spec.ForProvider.ProvisionedThroughput.WriteCapacityUnits)
		}
		res.SetProvisionedThroughput(f5)
	}

	return res
}

// GenerateUpdateTableInput returns an update input.
func GenerateUpdateTableInput(cr *svcapitypes.Table) *svcsdk.UpdateTableInput {
	res := &svcsdk.UpdateTableInput{}

	if cr.Spec.ForProvider.AttributeDefinitions != nil {
		f0 := []*svcsdk.AttributeDefinition{}
		for _, f0iter := range cr.Spec.ForProvider.AttributeDefinitions {
			f0elem := &svcsdk.AttributeDefinition{}
			if f0iter.AttributeName != nil {
				f0elem.SetAttributeName(*f0iter.AttributeName)
			}
			if f0iter.AttributeType != nil {
				f0elem.SetAttributeType(*f0iter.AttributeType)
			}
			f0 = append(f0, f0elem)
		}
		res.SetAttributeDefinitions(f0)
	}
	if cr.Spec.ForProvider.ProvisionedThroughput != nil {
		f3 := &svcsdk.ProvisionedThroughput{}
		if cr.Spec.ForProvider.ProvisionedThroughput.ReadCapacityUnits != nil {
			f3.SetReadCapacityUnits(*cr.Spec.ForProvider.ProvisionedThroughput.ReadCapacityUnits)
		}
		if cr.Spec.ForProvider.ProvisionedThroughput.WriteCapacityUnits != nil {
			f3.SetWriteCapacityUnits(*cr.Spec.ForProvider.ProvisionedThroughput.WriteCapacityUnits)
		}
		res.SetProvisionedThroughput(f3)
	}
	if cr.Status.AtProvider.TableName != nil {
		res.SetTableName(*cr.Status.AtProvider.TableName)
	}

	return res
}

// GenerateDeleteTableInput returns a deletion input.
func GenerateDeleteTableInput(cr *svcapitypes.Table) *svcsdk.DeleteTableInput {
	res := &svcsdk.DeleteTableInput{}
	res.TableName = aws.String(meta.GetExternalName(cr))
	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
