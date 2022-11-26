package utils

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/miladbonakdar/tp-rate-review/logger"
	"go.uber.org/zap"
)

func MarshalDynamoObject(obj interface{}) (map[string]*dynamodb.AttributeValue, error) {
	av, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		logger.New().Error("Error marshalling item",
			zap.String("trace", "MarshalDynamoObject"),
			zap.String("message", err.Error()))
		return nil, err
	}
	return av, nil
}
