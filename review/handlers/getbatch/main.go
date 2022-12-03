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

	req, err := dto.NewGetBatchReviews(request.Body)
	if err != nil {
		return utils.HandleFailOp(err)
	}

	if message := req.Validate(); message != "" {
		return utils.NewBadRequestRes(message), nil
	}

	repo := review.NewRepo()
	items, err := repo.GetBatchReviews(req.UserIds)

	if err != nil {
		return utils.HandleFailOp(err)
	}
	res := dto.GetUserReviewsDtoFromModel(items)
	return utils.NewCompleteResponse(res), nil
}

func main() {
	lambda.Start(Handler)
}
