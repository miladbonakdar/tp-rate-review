package models

type DefaultReviewModel struct {
	key         HashRange
	Rate        uint8  `json:"rate"`
	Review      string `json:"review"`
	Description string `json:"description,omitempty"`
}
