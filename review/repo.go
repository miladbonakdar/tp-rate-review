package review

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/miladbonakdar/tp-rate-review/db"
)

type reviewRepository struct {
	db *dynamodb.DynamoDB
}

func New() ReviewRepository {
	return &reviewRepository{
		db: db.NewDb(),
	}
}
