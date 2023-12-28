package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"lambda_function_1/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
