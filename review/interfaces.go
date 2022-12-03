package review

import "github.com/miladbonakdar/tp-rate-review/models"

type ReviewRepository interface {
	Add(review *ReviewModel) error
	Delete(key models.HashRange) error
	GetUserReviews(userId string) ([]*ReviewModel, error)
	GetBatchReviews(userIds []string) ([]*ReviewModel, error)
}
