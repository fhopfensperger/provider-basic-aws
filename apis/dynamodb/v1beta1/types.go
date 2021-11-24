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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

type ArchivalSummary struct {
	ArchivalBackupARN *string `json:"archivalBackupARN,omitempty"`

	ArchivalDateTime *metav1.Time `json:"archivalDateTime,omitempty"`

	ArchivalReason *string `json:"archivalReason,omitempty"`
}

type AttributeDefinition struct {
	AttributeName *string `json:"attributeName,omitempty"`

	AttributeType *string `json:"attributeType,omitempty"`
}

type BillingModeSummary struct {
	BillingMode *string `json:"billingMode,omitempty"`

	LastUpdateToPayPerRequestDateTime *metav1.Time `json:"lastUpdateToPayPerRequestDateTime,omitempty"`
}

type CreateGlobalSecondaryIndexAction struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
}

type CreateReplicaAction struct {
	RegionName *string `json:"regionName,omitempty"`
}

type CreateReplicationGroupMemberAction struct {
	GlobalSecondaryIndexes []*ReplicaGlobalSecondaryIndex `json:"globalSecondaryIndexes,omitempty"`

	KMSMasterKeyID *string `json:"kmsMasterKeyID,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`

	RegionName *string `json:"regionName,omitempty"`
}

type DeleteGlobalSecondaryIndexAction struct {
	IndexName *string `json:"indexName,omitempty"`
}

type DeleteReplicaAction struct {
	RegionName *string `json:"regionName,omitempty"`
}

type DeleteReplicationGroupMemberAction struct {
	RegionName *string `json:"regionName,omitempty"`
}

type GlobalSecondaryIndex struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
}

type KeySchemaElement struct {
	AttributeName *string `json:"attributeName,omitempty"`

	KeyType *string `json:"keyType,omitempty"`
}

type Projection struct {
	NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty"`

	ProjectionType *string `json:"projectionType,omitempty"`
}

type ProvisionedThroughput struct {
	ReadCapacityUnits *int64 `json:"readCapacityUnits,omitempty"`

	WriteCapacityUnits *int64 `json:"writeCapacityUnits,omitempty"`
}

type ProvisionedThroughputOverride struct {
	ReadCapacityUnits *int64 `json:"readCapacityUnits,omitempty"`
}

type Replica struct {
	RegionName *string `json:"regionName,omitempty"`
}

type ReplicaDescription struct {
	GlobalSecondaryIndexes []*ReplicaGlobalSecondaryIndexDescription `json:"globalSecondaryIndexes,omitempty"`

	KMSMasterKeyID *string `json:"kmsMasterKeyID,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`

	RegionName *string `json:"regionName,omitempty"`

	ReplicaInaccessibleDateTime *metav1.Time `json:"replicaInaccessibleDateTime,omitempty"`

	ReplicaStatus *string `json:"replicaStatus,omitempty"`

	ReplicaStatusDescription *string `json:"replicaStatusDescription,omitempty"`

	ReplicaStatusPercentProgress *string `json:"replicaStatusPercentProgress,omitempty"`
}

type ReplicaGlobalSecondaryIndex struct {
	IndexName *string `json:"indexName,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`
}

type ReplicaGlobalSecondaryIndexDescription struct {
	IndexName *string `json:"indexName,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`
}

type RestoreSummary struct {
	RestoreDateTime *metav1.Time `json:"restoreDateTime,omitempty"`

	RestoreInProgress *bool `json:"restoreInProgress,omitempty"`

	SourceBackupARN *string `json:"sourceBackupARN,omitempty"`

	SourceTableARN *string `json:"sourceTableARN,omitempty"`
}

type SSEDescription struct {
	InaccessibleEncryptionDateTime *metav1.Time `json:"inaccessibleEncryptionDateTime,omitempty"`

	KMSMasterKeyARN *string `json:"kmsMasterKeyARN,omitempty"`

	SSEType *string `json:"sseType,omitempty"`

	Status *string `json:"status,omitempty"`
}

type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}
