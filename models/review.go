package models

type ReviewModel struct {
	key     HashRange
	Rate    uint8  `json:"rate"`
	Review  string `json:"review"`
	From    string `json:"from"`
	Date    string `json:"date"`
	Session string `json:"session,omitempty"`
	Vote    uint8  `json:"vote"`
}
