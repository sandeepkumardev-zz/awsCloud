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
)

func GetItem(user *models.User) (res map[string]interface{}, err error) {
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
		if item["username"] == user.Username {
			return item, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func FindItem(user *models.User) (bool, error) {
	_, err := GetItem(user)
	if err != nil {
		switch {
		case err.Error() == "user not found":
			return true, nil
		default:
			return false, err
		}
	}
	return false, err
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
