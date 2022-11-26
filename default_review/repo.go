package default_review

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/miladbonakdar/tp-rate-review/db"
	"github.com/miladbonakdar/tp-rate-review/models"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

type defaultReviewRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func (drr *defaultReviewRepository) Add(review *DefaultReviewModel) error {
	obj, err := utils.MarshalDynamoObject(review)
	_, err = drr.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(drr.tableName),
		Item:      obj,
	})

	return err
}
func (drr *defaultReviewRepository) Delete(key models.HashRange) error {
	return nil
}
func (drr *defaultReviewRepository) GetDefaultReviews() ([]*DefaultReviewModel, error) {
	return nil, nil
}

func New() DefaultReviewRepository {
	return &defaultReviewRepository{
		db:        db.NewDb(),
		tableName: os.Getenv("DYNAMODB_TABLE"),
	}
}
