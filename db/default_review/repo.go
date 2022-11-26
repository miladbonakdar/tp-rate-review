package default_review

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/miladbonakdar/tp-rate-review/db"
)

type defaultReviewRepository struct {
	db *dynamodb.DynamoDB
}

func New() DefaultReviewRepository {
	return &defaultReviewRepository{
		db: db.NewDb(),
	}
}
