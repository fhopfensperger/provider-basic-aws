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

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ResourceCredentialsSecretRegionKey is the key for region that the S3 bucket is located
	ResourceCredentialsSecretRegionKey = "region"
)

// BucketParameters are parameters for configuring the calls made to AWS Bucket API.
type BucketParameters struct {

	// LocationConstraint specifies the Region where the bucket will be created.
	// It is a required field.
	LocationConstraint string `json:"locationConstraint"`
}

// BucketSpec represents the desired state of the Bucket.
type BucketSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BucketParameters `json:"forProvider"`
}

// BucketExternalStatus keeps the state for the external resource
type BucketExternalStatus struct {
	// ARN is the Amazon Resource Name (ARN) specifying the S3 Bucket. For more information
	// about ARNs and how to use them, see S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// in the Amazon Simple Storage Service guide.
	ARN string `json:"arn"`

	// Date the bucket was created. This date can change when making changes to
	// your bucket, such as editing its bucket policy.
	CreationDate *metav1.Time `json:"creationDate,omitempty"`
}

// BucketStatus represents the observed state of the Bucket.
type BucketStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BucketExternalStatus `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// An Bucket is a managed resource that represents an AWS S3 Bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
