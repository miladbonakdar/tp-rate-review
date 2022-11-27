package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miladbonakdar/tp-rate-review/defaultreview"
	"github.com/miladbonakdar/tp-rate-review/defaultreview/dto"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

func Handler(request events.APIGatewayProxyRequest) (utils.Response, error) {

	defaultReviewRepo := defaultreview.New()

	req, err := dto.NewCreateDefaultReview(request.Body)

	if err != nil {
		return utils.NewBadRequestRes(err.Error()), err
	}

	if message := req.Validate(); message != "" {
		return utils.NewBadRequestRes(message), errors.New(message)
	}

	model := req.ToDefaultReviewModel()
	err = defaultReviewRepo.Add(model)
	if err != nil {
		return utils.NewUnhandledEvent(err.Error()), err
	}
	resBody := dto.DefaultReviewDtoFromModel(model)
	return utils.NewCompleteResponse(resBody), nil
}

func main() {
	lambda.Start(Handler)
}
