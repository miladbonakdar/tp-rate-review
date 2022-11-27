package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/miladbonakdar/tp-rate-review/constants"
)

type HashRange struct {
	PK string `json:"pk"`
	SK string `json:"sk"`
}

func (h HashRange) DynamoKeyAttributes() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		constants.HashKey: {
			S: aws.String(h.PK),
		},
		constants.RangeKey: {
			S: aws.String(h.SK),
		},
	}
}
