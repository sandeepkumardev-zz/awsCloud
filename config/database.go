package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var DB_client *dynamodb.Client
var env = GetEnvVar()
var S3session *s3.S3

func ConnectionDB() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(env.AWS_REGION))
	if err != nil {
		return nil, err
	}

	S3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})))

	DB_client = dynamodb.NewFromConfig(cfg)

	// check if table exists
	resp, errr := GetTableInfo(DB_client, os.Getenv("TABLE_NAME"))
	if errr != nil {
		// create a new table
		fmt.Println("Creating new table : " + os.Getenv("TABLE_NAME"))
		_, err := createTable(DB_client, os.Getenv("TABLE_NAME"))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Table created!")
		}
	} else {
		fmt.Println("Table Size (bytes)", resp.Table.TableSizeBytes)
	}

	return DB_client, nil
}
