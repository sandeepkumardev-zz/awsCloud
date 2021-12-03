package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var DB_client *dynamodb.Client
var S3_client *s3.Client

func ConnectionDB() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return nil, err
	}

	S3_client = s3.NewFromConfig(cfg)
	DB_client = dynamodb.NewFromConfig(cfg)

	var tableName = os.Getenv("TABLE_NAME")
	// check if table exists
	resp, errr := GetTableInfo(DB_client, tableName)
	if errr != nil {
		// create a new table
		fmt.Println("Creating new table : " + tableName)
		_, err := createTable(DB_client, tableName)
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
