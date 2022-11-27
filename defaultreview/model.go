package defaultreview

import (
	"fmt"

	"github.com/miladbonakdar/tp-rate-review/models"
)

type DefaultReviewModel struct {
	Key         models.HashRange
	Rate        uint8  `json:"rate"`
	Review      string `json:"review"`
	Order       int    `json:"order"`
	Description string `json:"description,omitempty"`
}

func (d *DefaultReviewModel) LoadKeys() {
	d.Key.PK = fmt.Sprintf("default_rate#%v", d.Rate)
	d.Key.SK = fmt.Sprintf("review#%v-%s", d.Order, d.Review)
}
