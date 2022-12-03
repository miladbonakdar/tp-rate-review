package dto

import (
	"strings"

	"github.com/miladbonakdar/tp-rate-review/review"
)

type userReviewDetailMapItem struct {
	rateCount        int
	rateSum          int
	bestCustomReview *review.ReviewModel
}

type UserReviewDetail struct {
	BestCustomReview *ReviewDto `json:"bestCustomReview"`
	UserId           string     `json:"userId"`
	AverageRate      uint8      `json:"averageRate"`
}

type BatchReviewsDto struct {
	Users []*UserReviewDetail `json:"users"`
}

func BatchReviewsDtoFromModel(items []*review.ReviewModel) *BatchReviewsDto {
	length := len(items)
	res := &BatchReviewsDto{
		Users: make([]*UserReviewDetail, 0),
	}

	userReviews := make(map[string]*userReviewDetailMapItem)

	for i := 0; i < length; i++ {
		if items[i] == nil {
			continue
		}
		if _, ok := userReviews[items[i].User]; !ok {
			userReviews[items[i].User] = &userReviewDetailMapItem{}
		}
		userReviews[items[i].User].rateCount++
		userReviews[items[i].User].rateSum += int(items[i].Rate)

		if strings.Contains(items[i].SK, "custom") {
			userReviews[items[i].User].bestCustomReview = items[i]
		}
	}
	for userId, v := range userReviews {
		var averageRate uint8
		var bestReview *ReviewDto
		if v.rateCount >= 5 {
			averageRate = uint8(v.rateSum / v.rateCount)
		} else {
			averageRate = 5
		}
		if v.bestCustomReview != nil {
			bestReview = ReviewDtoFromModel(v.bestCustomReview)
		}

		res.Users = append(res.Users, &UserReviewDetail{
			UserId:           userId,
			AverageRate:      averageRate,
			BestCustomReview: bestReview,
		})

	}

	return res
}
