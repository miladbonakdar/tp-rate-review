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
	userId := request.PathParameters["userId"]

	if userId == "" {
		return utils.NewBadRequestRes("userId is not valid"), nil
	}

	repo := review.NewRepo()
	items, err := repo.GetUserReviews(userId)

	if err != nil {
		return utils.HandleFailOp(err)
	}
	res := dto.UserReviewsDtoFromModel(items)
	return utils.NewCompleteResponse(res), nil
}

func main() {
	lambda.Start(Handler)
}
