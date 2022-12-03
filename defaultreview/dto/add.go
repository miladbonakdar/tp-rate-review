package dto

import (
	"encoding/json"

	"github.com/miladbonakdar/tp-rate-review/defaultreview"
	"github.com/miladbonakdar/tp-rate-review/fail"
)

type CreateDefaultReview struct {
	Rate        uint8  `json:"rate"`
	Order       int    `json:"order"`
	Review      string `json:"review"`
	Description string `json:"description"`
}

func (c CreateDefaultReview) Validate() string {
	if c.Rate < 1 || c.Rate > 5 {
		return "Rate is not valid, [1,5] range"
	}
	if c.Review == "" {
		return "Review is empty"
	}
	return ""
}

func (c CreateDefaultReview) ToDefaultReviewModel() *defaultreview.DefaultReviewModel {
	model := &defaultreview.DefaultReviewModel{
		Rate:        c.Rate,
		Review:      c.Review,
		Description: c.Description,
		Order:       c.Order,
	}
	model.LoadKeys()
	return model
}

func NewCreateDefaultReview(reqBody string) (*CreateDefaultReview, error) {
	var req CreateDefaultReview
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "NewCreateDefaultReview")
	}
	return &req, nil
}
