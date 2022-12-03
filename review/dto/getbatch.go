package dto

import (
	"encoding/json"

	"github.com/miladbonakdar/tp-rate-review/fail"
)

type BatchReviews struct {
	UserIds []string `json:"userIds"`
}

func (c BatchReviews) Validate() string {
	if len(c.UserIds) == 0 {
		return "user ids list is empty"
	}
	return ""
}

func NewBatchReviews(reqBody string) (*BatchReviews, error) {
	var req BatchReviews
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "BatchReviews")
	}
	return &req, nil
}
