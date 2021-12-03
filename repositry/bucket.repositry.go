package repositry

import (
	"awsCloud/config"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3CreateBucketAPI interface {
	CreateBucket(ctx context.Context,
		params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

func MakeBucket(c context.Context, api S3CreateBucketAPI, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return api.CreateBucket(c, input)
}

func CreateBucket(bucketName string) error {
	var bucket *string = &bucketName

	input := &s3.CreateBucketInput{
		Bucket: bucket,
	}

	_, err := MakeBucket(context.TODO(), config.S3client, input)

	return err
}
