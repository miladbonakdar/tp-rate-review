package dto

import (
	"strings"

	"github.com/miladbonakdar/tp-rate-review/review"
)

type GetUserReviewsDto struct {
	Reviews       []*ReviewDto `json:"reviews"`
	CustomReviews []*ReviewDto `json:"customReviews"`
	UserId        string       `json:"userId"`
	AverageRate   uint8        `json:"averageRate"`
}

func GetUserReviewsDtoFromModel(items []*review.ReviewModel) *GetUserReviewsDto {
	length := len(items)
	res := &GetUserReviewsDto{
		Reviews:       make([]*ReviewDto, 0),
		CustomReviews: make([]*ReviewDto, 0),
	}
	count := 0
	var totalRateSum = 0
	for i := 0; i < length; i++ {
		if items[i] == nil {
			continue
		}
		count++
		totalRateSum += int(items[i].Rate)
		if strings.Contains(items[i].SK, "custom") {
			res.CustomReviews = append(res.CustomReviews, ReviewDtoFromModel(items[i]))
		} else {
			res.Reviews = append(res.Reviews, ReviewDtoFromModel(items[i]))
		}
	}
	if len(items) != 0 {
		res.UserId = items[0].User
	}

	if count >= 5 {
		res.AverageRate = uint8(totalRateSum / count)
	} else {
		res.AverageRate = 5
	}

	return res
}
