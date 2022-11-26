package db

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/miladbonakdar/tp-rate-review/session"
)

func NewDb() *dynamodb.DynamoDB {
	return dynamodb.New(session.New())
}
