package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miladbonakdar/tp-rate-review/models"
	"github.com/miladbonakdar/tp-rate-review/review"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

func Handler(request events.APIGatewayProxyRequest) (utils.Response, error) {
	utils.LogRequest(&request)
	paramPk := request.PathParameters["pk"]
	paramSk := request.PathParameters["sk"]

	key := models.HashRangeKey{
		PK: paramPk,
		SK: paramSk,
	}

	repo := review.NewRepo()

	err := repo.Delete(key)

	if err != nil {
		return utils.HandleFailOp(err)
	}

	return utils.NewCompleteTextResponse("review removed"), nil
}

func main() {
	lambda.Start(Handler)
}
