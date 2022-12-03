package dto

import (
	"encoding/json"

	"github.com/miladbonakdar/tp-rate-review/fail"
)

type GetBatchReviews struct {
	UserIds []string `json:"userIds"`
}

func (c GetBatchReviews) Validate() string {
	if len(c.UserIds) == 0 {
		return "user ids list is empty"
	}
	return ""
}

func NewGetBatchReviews(reqBody string) (*GetBatchReviews, error) {
	var req GetBatchReviews
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "NewGetBatchReviews")
	}
	return &req, nil
}
