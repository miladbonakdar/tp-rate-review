package defaultreview

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/miladbonakdar/tp-rate-review/logger"
	"github.com/miladbonakdar/tp-rate-review/models"
	"go.uber.org/zap"
)

type DefaultReviewModel struct {
	Key         models.HashRange
	Rate        uint8  `json:"rate"`
	Review      string `json:"review"`
	Order       int    `json:"order"`
	Description string `json:"description,omitempty"`
}

func (d *DefaultReviewModel) LoadKeys() {
	d.Key.PK = fmt.Sprintf(pkFormat, d.Rate)
	d.Key.SK = fmt.Sprintf(skFormat, d.Order, d.Review)
}

func NewDefaultReviewModel(item map[string]*dynamodb.AttributeValue) *DefaultReviewModel {
	if len(item) == 0 {
		return nil
	}
	var defaultReview DefaultReviewModel
	err := dynamodbattribute.UnmarshalMap(item, &defaultReview)
	if err != nil {
		logger.New().Error("error while trying to unmarshal map 'NewDefaultReviewModel'",
			zap.String("message", err.Error()))
		return nil
	}

	return &defaultReview
}

func NewDefaultReviewModelList(items []map[string]*dynamodb.AttributeValue) []*DefaultReviewModel {
	length := len(items)
	list := make([]*DefaultReviewModel, length)
	for i := 0; i < length; i++ {
		list[i] = NewDefaultReviewModel(items[i])
	}
	return list
}
