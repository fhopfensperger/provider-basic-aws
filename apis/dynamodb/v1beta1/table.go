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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TableParameters defines the desired state of Table
type TableParameters struct {
	// Region is which region the Table will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// An array of attributes that describe the key schema for the table and indexes.
	// +kubebuilder:validation:Required
	AttributeDefinitions []*AttributeDefinition `json:"attributeDefinitions"`

	// Specifies the attributes that make up the primary key for a table or an index.
	// The attributes in KeySchema must also be defined in the AttributeDefinitions
	// array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	// Each KeySchemaElement in the array is composed of:
	//
	//    * AttributeName - The name of this key attribute.
	//
	//    * KeyType - The role that the key attribute will assume: HASH - partition
	//    key RANGE - sort key
	//
	// The partition key of an item is also known as its hash attribute. The term
	// "hash attribute" derives from the DynamoDB usage of an internal hash function
	// to evenly distribute data items across partitions, based on their partition
	// key values.
	//
	// The sort key of an item is also known as its range attribute. The term "range
	// attribute" derives from the way DynamoDB stores items with the same partition
	// key physically close together, in sorted order by the sort key value.
	//
	// For a simple primary key (partition key), you must provide exactly one element
	// with a KeyType of HASH.
	//
	// For a composite primary key (partition key and sort key), you must provide
	// exactly two elements, in this order: The first element must have a KeyType
	// of HASH, and the second element must have a KeyType of RANGE.
	//
	// For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key)
	// in the Amazon DynamoDB Developer Guide.
	// +kubebuilder:validation:Required
	KeySchema []*KeySchemaElement `json:"keySchema"`

	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// If you set BillingMode as PROVISIONED, you must specify this property. If
	// you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`

	// A list of key-value pairs to label the table. For more information, see Tagging
	// for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).
	Tags []*Tag `json:"tags,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TableParameters `json:"forProvider"`
}

// TableObservation defines the observed state of Table
type TableObservation struct {
	// The date and time when the table was created, in UNIX epoch time (http://www.epochconverter.com/)
	// format.
	CreationDateTime *metav1.Time `json:"creationDateTime,omitempty"`
	// The Amazon Resource Name (ARN) that uniquely identifies the table.
	TableARN *string `json:"tableARN,omitempty"`
	// Unique identifier for the table for which the backup was created.
	TableID *string `json:"tableID,omitempty"`
	// The name of the table.
	TableName *string `json:"tableName,omitempty"`
	// The total size of the specified table, in bytes. DynamoDB updates this value
	// approximately every six hours. Recent changes might not be reflected in this
	// value.
	TableSizeBytes *int64 `json:"tableSizeBytes,omitempty"`
	// The current state of the table:
	//
	//    * CREATING - The table is being created.
	//
	//    * UPDATING - The table is being updated.
	//
	//    * DELETING - The table is being deleted.
	//
	//    * ACTIVE - The table is ready for use.
	//
	//    * INACCESSIBLE_ENCRYPTION_CREDENTIALS - The AWS KMS key used to encrypt
	//    the table in inaccessible. Table operations may fail due to failure to
	//    use the AWS KMS key. DynamoDB will initiate the table archival process
	//    when a table's AWS KMS key remains inaccessible for more than seven days.
	//
	//    * ARCHIVING - The table is being archived. Operations are not allowed
	//    until archival is complete.
	//
	//    * ARCHIVED - The table has been archived. See the ArchivalReason for more
	//    information.
	TableStatus *string `json:"tableStatus,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Table is the Schema for the Tables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	TableKind             = "Table"
	TableGroupKind        = schema.GroupKind{Group: Group, Kind: TableKind}.String()
	TableKindAPIVersion   = TableKind + "." + GroupVersion.String()
	TableGroupVersionKind = GroupVersion.WithKind(TableKind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
