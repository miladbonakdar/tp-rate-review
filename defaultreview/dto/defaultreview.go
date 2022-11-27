package dto

import (
	"github.com/miladbonakdar/tp-rate-review/defaultreview"
)

type DefaultReviewDto struct {
	PK          string `json:"pk"`
	SK          string `json:"sk"`
	Rate        uint8  `json:"rate"`
	Review      string `json:"review"`
	Order       int    `json:"order"`
	Description string `json:"description"`
}

func DefaultReviewDtoFromModel(r *defaultreview.DefaultReviewModel) *DefaultReviewDto {
	return &DefaultReviewDto{
		SK:          r.SK,
		PK:          r.PK,
		Rate:        r.Rate,
		Review:      r.Review,
		Order:       r.Order,
		Description: r.Description,
	}
}
