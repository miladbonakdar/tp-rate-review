package defaultreview

import "github.com/miladbonakdar/tp-rate-review/models"

type DefaultReviewRepository interface {
	Add(review *DefaultReviewModel) error
	Delete(key models.HashRange) error
	GetDefaultReviews() ([]*DefaultReviewModel, error)
}
