package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

func NewBadRequestRes(message string) Response {
	body, _ := json.Marshal(map[string]interface{}{
		"message": message,
	})
	return Response{
		StatusCode:      400,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers:         defaultHeaders,
	}
}

func NewUnhandledEvent(message string) Response {
	body, _ := json.Marshal(map[string]interface{}{
		"message": message,
	})

	return Response{
		StatusCode:      500,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers:         defaultHeaders,
	}
}

func NewCompleteResponse(item interface{}) Response {
	body, err := json.Marshal(item)
	if err != nil {
		return NewUnhandledEvent(err.Error())
	}
	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers:         defaultHeaders,
	}
}
