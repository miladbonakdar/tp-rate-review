package review

type ReviewModel struct {
	PK      string `json:"pk"`
	SK      string `json:"sk"`
	Rate    uint8  `json:"rate"`
	Review  string `json:"review"`
	From    string `json:"from"`
	Date    string `json:"date"`
	Session string `json:"session,omitempty"`
}
