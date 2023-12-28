package handler

import (
	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, lambdaError error) {
	result := "1"

	return events.APIGatewayProxyResponse{Body: result, StatusCode: 200}, nil
}
