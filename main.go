package main

import (
	"github.com/jamesyeap/lambda_dx"
	lambdahandler1 "lambda_function_1/handler"
	lambdahandler2 "lambda_function_2/handler"
)

func main() {
	server := lambda_dx.NewBoxHttpServer()

	/* add all the lambda functions for testing here */
	server.AddLambdaFunction("/route_to_lambda_function_1", lambdahandler1.HandleRequest)
	server.AddLambdaFunction("/route_to_lambda_function_2", lambdahandler2.HandleRequest)
	// ... ... ... ... ... ... ... ... ... ... ... ... ... ... ...
	// ... ... add more lambda functions for testing here ... ....
	// ... ... ... ... ... ... ... ... ... ... ... ... ... ... ...

	server.Start()
}
