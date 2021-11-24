//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchivalSummary) DeepCopyInto(out *ArchivalSummary) {
	*out = *in
	if in.ArchivalBackupARN != nil {
		in, out := &in.ArchivalBackupARN, &out.ArchivalBackupARN
		*out = new(string)
		**out = **in
	}
	if in.ArchivalDateTime != nil {
		in, out := &in.ArchivalDateTime, &out.ArchivalDateTime
		*out = (*in).DeepCopy()
	}
	if in.ArchivalReason != nil {
		in, out := &in.ArchivalReason, &out.ArchivalReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchivalSummary.
func (in *ArchivalSummary) DeepCopy() *ArchivalSummary {
	if in == nil {
		return nil
	}
	out := new(ArchivalSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributeDefinition) DeepCopyInto(out *AttributeDefinition) {
	*out = *in
	if in.AttributeName != nil {
		in, out := &in.AttributeName, &out.AttributeName
		*out = new(string)
		**out = **in
	}
	if in.AttributeType != nil {
		in, out := &in.AttributeType, &out.AttributeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributeDefinition.
func (in *AttributeDefinition) DeepCopy() *AttributeDefinition {
	if in == nil {
		return nil
	}
	out := new(AttributeDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingModeSummary) DeepCopyInto(out *BillingModeSummary) {
	*out = *in
	if in.BillingMode != nil {
		in, out := &in.BillingMode, &out.BillingMode
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateToPayPerRequestDateTime != nil {
		in, out := &in.LastUpdateToPayPerRequestDateTime, &out.LastUpdateToPayPerRequestDateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingModeSummary.
func (in *BillingModeSummary) DeepCopy() *BillingModeSummary {
	if in == nil {
		return nil
	}
	out := new(BillingModeSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateGlobalSecondaryIndexAction) DeepCopyInto(out *CreateGlobalSecondaryIndexAction) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]*KeySchemaElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeySchemaElement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Projection != nil {
		in, out := &in.Projection, &out.Projection
		*out = new(Projection)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisionedThroughput != nil {
		in, out := &in.ProvisionedThroughput, &out.ProvisionedThroughput
		*out = new(ProvisionedThroughput)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateGlobalSecondaryIndexAction.
func (in *CreateGlobalSecondaryIndexAction) DeepCopy() *CreateGlobalSecondaryIndexAction {
	if in == nil {
		return nil
	}
	out := new(CreateGlobalSecondaryIndexAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateReplicaAction) DeepCopyInto(out *CreateReplicaAction) {
	*out = *in
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateReplicaAction.
func (in *CreateReplicaAction) DeepCopy() *CreateReplicaAction {
	if in == nil {
		return nil
	}
	out := new(CreateReplicaAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateReplicationGroupMemberAction) DeepCopyInto(out *CreateReplicationGroupMemberAction) {
	*out = *in
	if in.GlobalSecondaryIndexes != nil {
		in, out := &in.GlobalSecondaryIndexes, &out.GlobalSecondaryIndexes
		*out = make([]*ReplicaGlobalSecondaryIndex, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicaGlobalSecondaryIndex)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KMSMasterKeyID != nil {
		in, out := &in.KMSMasterKeyID, &out.KMSMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.ProvisionedThroughputOverride != nil {
		in, out := &in.ProvisionedThroughputOverride, &out.ProvisionedThroughputOverride
		*out = new(ProvisionedThroughputOverride)
		(*in).DeepCopyInto(*out)
	}
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateReplicationGroupMemberAction.
func (in *CreateReplicationGroupMemberAction) DeepCopy() *CreateReplicationGroupMemberAction {
	if in == nil {
		return nil
	}
	out := new(CreateReplicationGroupMemberAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteGlobalSecondaryIndexAction) DeepCopyInto(out *DeleteGlobalSecondaryIndexAction) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteGlobalSecondaryIndexAction.
func (in *DeleteGlobalSecondaryIndexAction) DeepCopy() *DeleteGlobalSecondaryIndexAction {
	if in == nil {
		return nil
	}
	out := new(DeleteGlobalSecondaryIndexAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteReplicaAction) DeepCopyInto(out *DeleteReplicaAction) {
	*out = *in
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteReplicaAction.
func (in *DeleteReplicaAction) DeepCopy() *DeleteReplicaAction {
	if in == nil {
		return nil
	}
	out := new(DeleteReplicaAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteReplicationGroupMemberAction) DeepCopyInto(out *DeleteReplicationGroupMemberAction) {
	*out = *in
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteReplicationGroupMemberAction.
func (in *DeleteReplicationGroupMemberAction) DeepCopy() *DeleteReplicationGroupMemberAction {
	if in == nil {
		return nil
	}
	out := new(DeleteReplicationGroupMemberAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalSecondaryIndex) DeepCopyInto(out *GlobalSecondaryIndex) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]*KeySchemaElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeySchemaElement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Projection != nil {
		in, out := &in.Projection, &out.Projection
		*out = new(Projection)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisionedThroughput != nil {
		in, out := &in.ProvisionedThroughput, &out.ProvisionedThroughput
		*out = new(ProvisionedThroughput)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalSecondaryIndex.
func (in *GlobalSecondaryIndex) DeepCopy() *GlobalSecondaryIndex {
	if in == nil {
		return nil
	}
	out := new(GlobalSecondaryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeySchemaElement) DeepCopyInto(out *KeySchemaElement) {
	*out = *in
	if in.AttributeName != nil {
		in, out := &in.AttributeName, &out.AttributeName
		*out = new(string)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeySchemaElement.
func (in *KeySchemaElement) DeepCopy() *KeySchemaElement {
	if in == nil {
		return nil
	}
	out := new(KeySchemaElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Projection) DeepCopyInto(out *Projection) {
	*out = *in
	if in.NonKeyAttributes != nil {
		in, out := &in.NonKeyAttributes, &out.NonKeyAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProjectionType != nil {
		in, out := &in.ProjectionType, &out.ProjectionType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Projection.
func (in *Projection) DeepCopy() *Projection {
	if in == nil {
		return nil
	}
	out := new(Projection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedThroughput) DeepCopyInto(out *ProvisionedThroughput) {
	*out = *in
	if in.ReadCapacityUnits != nil {
		in, out := &in.ReadCapacityUnits, &out.ReadCapacityUnits
		*out = new(int64)
		**out = **in
	}
	if in.WriteCapacityUnits != nil {
		in, out := &in.WriteCapacityUnits, &out.WriteCapacityUnits
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedThroughput.
func (in *ProvisionedThroughput) DeepCopy() *ProvisionedThroughput {
	if in == nil {
		return nil
	}
	out := new(ProvisionedThroughput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedThroughputOverride) DeepCopyInto(out *ProvisionedThroughputOverride) {
	*out = *in
	if in.ReadCapacityUnits != nil {
		in, out := &in.ReadCapacityUnits, &out.ReadCapacityUnits
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedThroughputOverride.
func (in *ProvisionedThroughputOverride) DeepCopy() *ProvisionedThroughputOverride {
	if in == nil {
		return nil
	}
	out := new(ProvisionedThroughputOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Replica) DeepCopyInto(out *Replica) {
	*out = *in
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Replica.
func (in *Replica) DeepCopy() *Replica {
	if in == nil {
		return nil
	}
	out := new(Replica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaDescription) DeepCopyInto(out *ReplicaDescription) {
	*out = *in
	if in.GlobalSecondaryIndexes != nil {
		in, out := &in.GlobalSecondaryIndexes, &out.GlobalSecondaryIndexes
		*out = make([]*ReplicaGlobalSecondaryIndexDescription, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicaGlobalSecondaryIndexDescription)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KMSMasterKeyID != nil {
		in, out := &in.KMSMasterKeyID, &out.KMSMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.ProvisionedThroughputOverride != nil {
		in, out := &in.ProvisionedThroughputOverride, &out.ProvisionedThroughputOverride
		*out = new(ProvisionedThroughputOverride)
		(*in).DeepCopyInto(*out)
	}
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
	if in.ReplicaInaccessibleDateTime != nil {
		in, out := &in.ReplicaInaccessibleDateTime, &out.ReplicaInaccessibleDateTime
		*out = (*in).DeepCopy()
	}
	if in.ReplicaStatus != nil {
		in, out := &in.ReplicaStatus, &out.ReplicaStatus
		*out = new(string)
		**out = **in
	}
	if in.ReplicaStatusDescription != nil {
		in, out := &in.ReplicaStatusDescription, &out.ReplicaStatusDescription
		*out = new(string)
		**out = **in
	}
	if in.ReplicaStatusPercentProgress != nil {
		in, out := &in.ReplicaStatusPercentProgress, &out.ReplicaStatusPercentProgress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaDescription.
func (in *ReplicaDescription) DeepCopy() *ReplicaDescription {
	if in == nil {
		return nil
	}
	out := new(ReplicaDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaGlobalSecondaryIndex) DeepCopyInto(out *ReplicaGlobalSecondaryIndex) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
	if in.ProvisionedThroughputOverride != nil {
		in, out := &in.ProvisionedThroughputOverride, &out.ProvisionedThroughputOverride
		*out = new(ProvisionedThroughputOverride)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaGlobalSecondaryIndex.
func (in *ReplicaGlobalSecondaryIndex) DeepCopy() *ReplicaGlobalSecondaryIndex {
	if in == nil {
		return nil
	}
	out := new(ReplicaGlobalSecondaryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaGlobalSecondaryIndexDescription) DeepCopyInto(out *ReplicaGlobalSecondaryIndexDescription) {
	*out = *in
	if in.IndexName != nil {
		in, out := &in.IndexName, &out.IndexName
		*out = new(string)
		**out = **in
	}
	if in.ProvisionedThroughputOverride != nil {
		in, out := &in.ProvisionedThroughputOverride, &out.ProvisionedThroughputOverride
		*out = new(ProvisionedThroughputOverride)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaGlobalSecondaryIndexDescription.
func (in *ReplicaGlobalSecondaryIndexDescription) DeepCopy() *ReplicaGlobalSecondaryIndexDescription {
	if in == nil {
		return nil
	}
	out := new(ReplicaGlobalSecondaryIndexDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSummary) DeepCopyInto(out *RestoreSummary) {
	*out = *in
	if in.RestoreDateTime != nil {
		in, out := &in.RestoreDateTime, &out.RestoreDateTime
		*out = (*in).DeepCopy()
	}
	if in.RestoreInProgress != nil {
		in, out := &in.RestoreInProgress, &out.RestoreInProgress
		*out = new(bool)
		**out = **in
	}
	if in.SourceBackupARN != nil {
		in, out := &in.SourceBackupARN, &out.SourceBackupARN
		*out = new(string)
		**out = **in
	}
	if in.SourceTableARN != nil {
		in, out := &in.SourceTableARN, &out.SourceTableARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSummary.
func (in *RestoreSummary) DeepCopy() *RestoreSummary {
	if in == nil {
		return nil
	}
	out := new(RestoreSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSEDescription) DeepCopyInto(out *SSEDescription) {
	*out = *in
	if in.InaccessibleEncryptionDateTime != nil {
		in, out := &in.InaccessibleEncryptionDateTime, &out.InaccessibleEncryptionDateTime
		*out = (*in).DeepCopy()
	}
	if in.KMSMasterKeyARN != nil {
		in, out := &in.KMSMasterKeyARN, &out.KMSMasterKeyARN
		*out = new(string)
		**out = **in
	}
	if in.SSEType != nil {
		in, out := &in.SSEType, &out.SSEType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSEDescription.
func (in *SSEDescription) DeepCopy() *SSEDescription {
	if in == nil {
		return nil
	}
	out := new(SSEDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableObservation) DeepCopyInto(out *TableObservation) {
	*out = *in
	if in.CreationDateTime != nil {
		in, out := &in.CreationDateTime, &out.CreationDateTime
		*out = (*in).DeepCopy()
	}
	if in.TableARN != nil {
		in, out := &in.TableARN, &out.TableARN
		*out = new(string)
		**out = **in
	}
	if in.TableID != nil {
		in, out := &in.TableID, &out.TableID
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.TableSizeBytes != nil {
		in, out := &in.TableSizeBytes, &out.TableSizeBytes
		*out = new(int64)
		**out = **in
	}
	if in.TableStatus != nil {
		in, out := &in.TableStatus, &out.TableStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableObservation.
func (in *TableObservation) DeepCopy() *TableObservation {
	if in == nil {
		return nil
	}
	out := new(TableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.AttributeDefinitions != nil {
		in, out := &in.AttributeDefinitions, &out.AttributeDefinitions
		*out = make([]*AttributeDefinition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AttributeDefinition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeySchema != nil {
		in, out := &in.KeySchema, &out.KeySchema
		*out = make([]*KeySchemaElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeySchemaElement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ProvisionedThroughput != nil {
		in, out := &in.ProvisionedThroughput, &out.ProvisionedThroughput
		*out = new(ProvisionedThroughput)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}