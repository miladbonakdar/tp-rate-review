package default_review

import (
	"github.com/miladbonakdar/tp-rate-review/models"
)

type DefaultReviewModel struct {
	key         models.HashRange
	Rate        uint8  `json:"rate"`
	Review      string `json:"review"`
	Description string `json:"description,omitempty"`
}
