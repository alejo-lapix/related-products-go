package groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

var tableName = "groups"

func TestDynamoDBGroupRepository_Store(t *testing.T) {
	type args struct {
		product *Group
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Stores a new record",
			wantErr: false,
			args:    args{product: exampleGroup()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBGroupRepository{
				DynamoDB:  dynamoDBInstance(),
				tableName: &tableName,
			}
			if err := repository.Store(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBGroupRepository_FindByProduct(t *testing.T) {
	type args struct {
		primaryProductID *string
	}
	tests := []struct {
		name    string
		args    args
		want    *Group
		wantErr bool
	}{
		{
			name: "Find the created item",
			args: args{
				primaryProductID: exampleGroup().PrimaryProductID,
			},
			want:    exampleGroup(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBGroupRepository{
				DynamoDB:  dynamoDBInstance(),
				tableName: &tableName,
			}
			got, err := repository.FindByProduct(tt.args.primaryProductID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBGroupRepository_Remove(t *testing.T) {
	type args struct {
		ID *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Remove an item",
			args:    args{ID: exampleGroup().PrimaryProductID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBGroupRepository{
				DynamoDB:  dynamoDBInstance(),
				tableName: &tableName,
			}
			if err := repository.Remove(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func dynamoDBInstance() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		panic(err.Error())
	}

	return dynamodb.New(sess)
}

func exampleGroup() *Group {
	return &Group{
		PrimaryProductID: aws.String("abc"),
		Associations: []*Association{
			{
				ProductID: aws.String("def"),
				Ratio:     aws.Float64(1.5),
			},
			{
				ProductID: aws.String("ghi"),
				Ratio:     aws.Float64(0.5),
			},
		},
	}
}
