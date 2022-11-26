package review

import "github.com/miladbonakdar/tp-rate-review/models"

type ReviewModel struct {
	key     models.HashRange
	Rate    uint8  `json:"rate"`
	Review  string `json:"review"`
	From    string `json:"from"`
	Date    string `json:"date"`
	Session string `json:"session,omitempty"`
	Vote    uint8  `json:"vote"`
}
