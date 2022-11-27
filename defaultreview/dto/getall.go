package dto

import "github.com/miladbonakdar/tp-rate-review/defaultreview"

type GetAllDefaultReviewsDto struct {
	Items []*DefaultReviewDto `json:"items"`
	Rate  uint8               `json:"rate"`
}

func GetAllDefaultReviewsDtoFromModel(items []*defaultreview.DefaultReviewModel) *GetAllDefaultReviewsDto {
	length := len(items)
	res := &GetAllDefaultReviewsDto{
		Items: make([]*DefaultReviewDto, 0),
	}
	for i := 0; i < length; i++ {
		if items[i] == nil {
			continue
		}
		res.Items = append(res.Items, DefaultReviewDtoFromModel(items[i]))
	}
	if len(res.Items) != 0 {
		res.Rate = res.Items[0].Rate
	}
	return res
}
