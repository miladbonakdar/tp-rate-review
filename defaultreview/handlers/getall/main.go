package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miladbonakdar/tp-rate-review/defaultreview"
	"github.com/miladbonakdar/tp-rate-review/defaultreview/dto"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

func Handler(request events.APIGatewayProxyRequest) (utils.Response, error) {
	paramRate := request.PathParameters["rate"]
	val, err := utils.ParseToUint8(paramRate)
	if err != nil {
		return utils.NewBadRequestRes(err.Error()), nil
	}

	if val < 1 || val > 5 {
		errorMessage := "Rate is not valid, [1,5] range"
		return utils.NewBadRequestRes(errorMessage), nil
	}

	repo := defaultreview.NewRepo()
	items, err := repo.GetDefaultReviews(val)

	if err != nil {
		return utils.NewUnhandledEvent(err.Error()), err
	}
	res := dto.GetAllDefaultReviewsDtoFromModel(items)
	return utils.NewCompleteResponse(res), nil
}

func main() {
	lambda.Start(Handler)
}
