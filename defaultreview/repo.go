package defaultreview

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

type defaultReviewRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func (drr *defaultReviewRepository) Add(review *DefaultReviewModel) error {
	obj, err := utils.MarshalDynamoObject(review)
	if err != nil {
		return fail.NewFailByError(400, err, "Add defaultReviewRepository MarshalDynamoObject")
	}
	_, err = drr.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(drr.tableName),
		Item:      obj,
	})
	if err != nil {
		return fail.NewFailByError(400, err, "Add defaultReviewRepository PutItem")
	}

	return nil
}

func (drr *defaultReviewRepository) Delete(key models.HashRange) error {

	input := &dynamodb.DeleteItemInput{
		Key:       key.DynamoKeyAttributes(),
		TableName: aws.String(drr.tableName),
	}

	_, err := drr.db.DeleteItem(input)
	if err != nil {
		return fail.NewFailByError(400, err, "Delete defaultReviewRepository")
	}
	return nil
}

func (drr *defaultReviewRepository) GetDefaultReviews(rate uint8) ([]*DefaultReviewModel, error) {
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(drr.tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			constants.HashKey: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(fmt.Sprintf(pkFormat, rate)),
					},
				},
			},
		},
	}
	res, err := drr.db.Query(queryInput)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "GetDefaultReviews defaultReviewRepository")
	}

	if res.Count == aws.Int64(0) {
		return nil, nil
	}

	return NewDefaultReviewModelList(res.Items), nil
}

func NewRepo() DefaultReviewRepository {
	return &defaultReviewRepository{
		db:        db.NewDb(),
		tableName: os.Getenv("DYNAMODB_TABLE"),
	}
}
