package utils

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/miladbonakdar/tp-rate-review/logger"
	"go.uber.org/zap"
)

func LogRequest(request *events.APIGatewayProxyRequest) {
	logger.New().Info("request-log", zap.String("path", request.Path),
		zap.Any("PathParameters", request.PathParameters),
		zap.String("body", request.Body))
}
