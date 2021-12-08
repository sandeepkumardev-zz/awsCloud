package repositry

import (
	"awsCloud/config"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket(bucketName string) (resp *s3.CreateBucketOutput, err error) {
	resp, err = config.S3session.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(os.Getenv("AWS_REGION")),
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				return nil, fmt.Errorf("bucket name is already in use")

			case s3.ErrCodeBucketAlreadyOwnedByYou:
				return nil, fmt.Errorf("bucket exists and is owned by you")

			default:
				return nil, err
			}
		}
	}
	return resp, nil
}
