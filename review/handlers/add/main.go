package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miladbonakdar/tp-rate-review/review"
	"github.com/miladbonakdar/tp-rate-review/review/dto"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

func Handler(request events.APIGatewayProxyRequest) (utils.Response, error) {
	utils.LogRequest(&request)
	reviewRepo := review.NewRepo()

	req, err := dto.NewCreateReview(request.Body)
	if err != nil {
		return utils.HandleFailOp(err)
	}

	if message := req.Validate(); message != "" {
		return utils.NewBadRequestRes(message), nil
	}

	model := req.ToReviewModel()
	err = reviewRepo.Add(model)
	if err != nil {
		return utils.HandleFailOp(err)
	}
	resBody := dto.ReviewDtoFromModel(model)
	return utils.NewCompleteResponse(resBody), nil
}

func main() {
	lambda.Start(Handler)
}
