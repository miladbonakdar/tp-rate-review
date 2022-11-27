package dto

import (
	"github.com/miladbonakdar/tp-rate-review/defaultreview"
	"github.com/miladbonakdar/tp-rate-review/models"
)

type DefaultReviewDto struct {
	Key         models.HashRange
	Rate        uint8  `json:"rate"`
	Review      string `json:"review"`
	Order       int    `json:"order"`
	Description string `json:"description,omitempty"`
}

func DefaultReviewDtoFromModel(r *defaultreview.DefaultReviewModel) *DefaultReviewDto {
	return &DefaultReviewDto{
		Key:         r.Key,
		Rate:        r.Rate,
		Review:      r.Review,
		Order:       r.Order,
		Description: r.Description,
	}
}
