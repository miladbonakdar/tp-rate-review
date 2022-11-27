package dto

import (
	"encoding/json"

	"github.com/miladbonakdar/tp-rate-review/defaultreview"
)

type CreateDefaultReview struct {
	PK          string `json:"pk"`
	SK          string `json:"sk"`
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
		PK:          c.PK,
		SK:          c.SK,
	}
	model.LoadKeys()
	return model
}

func NewCreateDefaultReview(reqBody string) (*CreateDefaultReview, error) {
	var req CreateDefaultReview
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
