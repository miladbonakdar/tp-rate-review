package review

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/miladbonakdar/tp-rate-review/logger"
	"go.uber.org/zap"
)

type ReviewModel struct {
	PK      string `json:"pk"`
	SK      string `json:"sk"`
	Rate    uint8  `json:"rate"`
	Review  string `json:"review"`
	User    string `json:"user"`
	From    string `json:"from"`
	Date    string `json:"date"`
	Session string `json:"session,omitempty"`
}

// rate-review item :
//     hashkey : review#{user-ulid}
//     sort key : custom#{iso time stamp}
//     sort key : review#{iso time stamp}
//         rate: 1-5
//         review: string
//         from: user-ulid
//         date: iso date string
//         session: {session-id} can be null
//
// add user text review //custom
// add user default review // review
// get user reviews

func (d *ReviewModel) LoadReviewKeys() {
	d.PK = fmt.Sprintf(pkFormat, d.User)
	d.SK = fmt.Sprintf(skReviewFormat, time.Now().Unix())
}
func (d *ReviewModel) LoadCustomKeys() {
	d.PK = fmt.Sprintf(pkFormat, d.User)
	d.SK = fmt.Sprintf(skCustomFormat, time.Now().Unix())
}

func NewReviewModel(item map[string]*dynamodb.AttributeValue) *ReviewModel {
	if len(item) == 0 {
		return nil
	}
	var review ReviewModel
	err := dynamodbattribute.UnmarshalMap(item, &review)
	if err != nil {
		logger.New().Error("error while trying to unmarshal map 'NewReviewModel'",
			zap.String("message", err.Error()))
		return nil
	}

	return &review
}

func NewReviewModelList(items []map[string]*dynamodb.AttributeValue) []*ReviewModel {
	length := len(items)
	list := make([]*ReviewModel, length)
	for i := 0; i < length; i++ {
		list[i] = NewReviewModel(items[i])
	}
	return list
}
