package dto

import (
	"encoding/json"
	"time"

	"github.com/miladbonakdar/tp-rate-review/fail"
	"github.com/miladbonakdar/tp-rate-review/review"
)

type CreateReview struct {
	Rate    uint8  `json:"rate"`
	Review  string `json:"review"`
	Type    string `json:"type"`
	User    string `json:"user"`
	From    string `json:"from"`
	Session string `json:"session,omitempty"`
}

func (c CreateReview) Validate() string {
	if c.Rate < 1 || c.Rate > 5 {
		return "Rate is not valid, [1,5] range"
	}
	if c.Review == "" {
		return "Review is empty"
	}
	if c.Type == "" || (c.Type != review.TextReviewType && c.Type != review.DefaultReviewType) {
		return "Review type is not valid 'text' , 'default'"
	}
	if c.From == "" {
		return "From is empty"
	}
	if c.User == "" {
		return "User is empty"
	}
	return ""
}

func NewCreateReview(reqBody string) (*CreateReview, error) {
	var req CreateReview
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		return nil, fail.NewFailByError(400, err, "NewCreateReview")
	}
	return &req, nil
}

func (c *CreateReview) ToReviewModel() *review.ReviewModel {
	item := &review.ReviewModel{
		Rate:    c.Rate,
		Review:  c.Review,
		From:    c.From,
		User:    c.User,
		Date:    time.Now().Format(time.RFC3339),
		Session: c.Session,
	}
	if c.Type == review.TextReviewType {
		item.LoadCustomKeys()
	} else {
		item.LoadReviewKeys()
	}
	return item
}
