package dto

import "github.com/miladbonakdar/tp-rate-review/review"

type ReviewDto struct {
	PK      string `json:"pk"`
	SK      string `json:"sk"`
	Rate    uint8  `json:"rate"`
	Review  string `json:"review"`
	User    string `json:"user"`
	From    string `json:"from"`
	Date    string `json:"date"`
	Session string `json:"session,omitempty"`
}

func ReviewDtoFromModel(r *review.ReviewModel) *ReviewDto {
	return &ReviewDto{
		SK:      r.SK,
		PK:      r.PK,
		Rate:    r.Rate,
		Review:  r.Review,
		User:    r.User,
		From:    r.From,
		Date:    r.Date,
		Session: r.Session,
	}
}
