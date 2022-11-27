package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miladbonakdar/tp-rate-review/defaultreview"
	"github.com/miladbonakdar/tp-rate-review/models"
	"github.com/miladbonakdar/tp-rate-review/utils"
)

func Handler(request events.APIGatewayProxyRequest) (utils.Response, error) {
	paramPk := request.PathParameters["pk"]
	paramSk := request.PathParameters["sk"]

	key := models.HashRangeKey{
		PK: paramPk,
		SK: paramSk,
	}

	repo := defaultreview.NewRepo()

	err := repo.Delete(key)

	if err != nil {
		return utils.NewUnhandledEvent(err.Error()), err
	}
	return utils.NewCompleteTextResponse("default review removed"), nil
}

func main() {
	lambda.Start(Handler)
}
