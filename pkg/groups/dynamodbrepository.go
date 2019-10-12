package groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBGroupRepository struct {
	DynamoDB  *dynamodb.DynamoDB
	tableName *string
}

func NewDynamoDBGroupRepository(db *dynamodb.DynamoDB) *DynamoDBGroupRepository {
	return &DynamoDBGroupRepository{
		DynamoDB:  db,
		tableName: aws.String("groups"),
	}
}

func (repository *DynamoDBGroupRepository) Store(product *Group) error {
	item, err := dynamodbattribute.MarshalMap(product)

	if err != nil {
		return err
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: repository.tableName,
	})

	return err
}

func (repository *DynamoDBGroupRepository) FindByProduct(primaryProductID *string) (*Group, error) {
	output, err := repository.DynamoDB.GetItem(&dynamodb.GetItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"primaryProductId": {S: primaryProductID}},
		TableName: repository.tableName,
	})

	if err != nil || output.Item == nil {
		return nil, err
	}

	item := &Group{}
	err = dynamodbattribute.UnmarshalMap(output.Item, item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (repository *DynamoDBGroupRepository) Remove(ID *string) error {
	_, err := repository.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"primaryProductId": {S: ID}},
		TableName: repository.tableName,
	})

	return err
}
