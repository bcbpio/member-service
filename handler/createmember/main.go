package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bcbpio/member-service/handler"
)

func main() {
	//Register handler
	lambda.Start(handler.CreateMemberHandler)
}
