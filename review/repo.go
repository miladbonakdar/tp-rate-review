package review

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/miladbonakdar/tp-rate-review/constants"
	"github.com/miladbonakdar/tp-rate-review/db"
	"github.com/miladbonakdar/tp-rate-review/fail"
	"github.com/miladbonakdar/tp-rate-review/models"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

type reviewRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func (drr *reviewRepository) Add(review *ReviewModel) error {
	obj, err := utils.MarshalDynamoObject(review)
	if err != nil {
		return fail.NewFailByError(400, err, "Add reviewRepository MarshalDynamoObject")
	}
	_, err = drr.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(drr.tableName),
		Item:      obj,
	})
	if err != nil {
		return fail.NewFailByError(400, err, "Add reviewRepository PutItem")
	}

	return nil
}

func (drr *reviewRepository) Delete(key models.HashRange) error {

	input := &dynamodb.DeleteItemInput{
		Key:       key.DynamoKeyAttributes(),
		TableName: aws.String(drr.tableName),
	}

	_, err := drr.db.DeleteItem(input)
	return fail.NewFailByError(400, err, "Delete reviewRepository")
}

func (drr *reviewRepository) GetUserReviews(userId string) ([]*ReviewModel, error) {
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(drr.tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			constants.HashKey: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(fmt.Sprintf(pkFormat, userId)),
					},
				},
			},
		},
	}
	res, err := drr.db.Query(queryInput)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "GetUserReviews reviewRepository")
	}

	if res.Count == aws.Int64(0) {
		return nil, nil
	}

	return NewReviewModelList(res.Items), nil
}

func (drr *reviewRepository) GetBatchReviews(userIds []string) ([]*ReviewModel, error) {
	pkAttributes := make([]*dynamodb.AttributeValue, len(userIds))
	for i, id := range userIds {
		pkAttributes[i] = &dynamodb.AttributeValue{
			S: aws.String(fmt.Sprintf(pkFormat, id)),
		}
	}
	var scanInput = &dynamodb.ScanInput{
		TableName: aws.String(drr.tableName),
		ScanFilter: map[string]*dynamodb.Condition{
			constants.HashKey: {
				ComparisonOperator: aws.String("IN"),
				AttributeValueList: pkAttributes,
			},
		},
	}
	res, err := drr.db.Scan(scanInput)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "GetBatchReviews reviewRepository")
	}

	if res.Count == aws.Int64(0) {
		return nil, nil
	}

	return NewReviewModelList(res.Items), nil
}

func NewRepo() ReviewRepository {
	return &reviewRepository{
		db:        db.NewDb(),
		tableName: os.Getenv("DYNAMODB_TABLE"),
	}
}
