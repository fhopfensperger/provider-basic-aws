package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/fhopfensperger/provider-basic-aws/apis/s3/v1beta1"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IsNotFound helper function to test for NotFound error
func IsNotFound(err error) bool {
	var notFoundError *s3types.NotFound
	return errors.As(err, &notFoundError)
}

// IsAlreadyExists helper function to test for ErrCodeBucketAlreadyOwnedByYou error
func IsAlreadyExists(err error) bool {
	var alreadyOwnedByYou *s3types.BucketAlreadyOwnedByYou
	return errors.As(err, &alreadyOwnedByYou)
}

// GenerateCreateBucketInput creates the input for CreateBucket S3 Client request
func GenerateCreateBucketInput(name string, s v1beta1.BucketParameters) *s3.CreateBucketInput {
	cbi := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}
	if s.LocationConstraint != "us-east-1" {
		cbi.CreateBucketConfiguration = &s3types.CreateBucketConfiguration{LocationConstraint: s3types.BucketLocationConstraint(s.LocationConstraint)}
	}
	return cbi
}

// GenerateBucketObservation generates the ARN string and the creation time stamp for the external status
func GenerateBucketObservation(name string, buckets []s3types.Bucket) v1beta1.BucketExternalStatus {
	bes := v1beta1.BucketExternalStatus{}
	for _, bucket := range buckets {
		if aws.ToString(bucket.Name) == name {
			t := metav1.NewTime(*bucket.CreationDate)
			bes.CreationDate = &t
		}
	}
	bes.ARN = fmt.Sprintf("arn:aws:s3:::%s", name)
	return bes
}
