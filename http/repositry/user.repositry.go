package repositry

import (
	"awsCloud/database/config"
	"awsCloud/database/models"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetItem(username string) (res map[string]interface{}, err error) {
	var resp []map[string]interface{}

	out, err := config.DB_client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &resp)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	for _, item := range resp {
		if item["username"] == username {
			return item, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func FindItem(username string) (bool, error) {
	_, err := GetItem(username)
	if err != nil {
		switch {
		case err.Error() == "user not found":
			return false, nil
		default:
			return false, err
		}
	}
	return true, err
}

func PutItem(user *models.User) error {
	data, err := attributevalue.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("something went wrong")
	}

	_, dberr := config.DB_client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item:      data,
	})

	return dberr
}

func UpadteItem(id string, attValue string) error {
	_, dberr := config.DB_client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("set verified = :verified"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":verified": &types.AttributeValueMemberS{Value: attValue},
		},
	})

	return dberr
}
