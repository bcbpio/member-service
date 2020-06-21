package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bcbpio/member-service/handler"
	"github.com/bcbpio/member-service/repository"
	"github.com/bcbpio/member-service/service"
)

//Handler - handler for create member
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Connect to db
	driver, session, err := handler.Connect()
	if err != nil {
		return generateErrorResponse(err.Error(), 500), err
	}
	defer driver.Close()
	defer session.Close()

	//Get new contact service
	var memberSvc service.Service
	{
		repo := repository.NewRepository(session)
		memberSvc = service.NewService(repo)
	}

	//Unmarshal post request body
	var m repository.Member
	err = json.Unmarshal([]byte(req.Body), &m)
	if err != nil {
		fmt.Println(err)
		return generateErrorResponse(err.Error(), 500), nil
	}

	//Create member using post request body as query parameter
	memberID, err := memberSvc.CreateMember(m)
	if err != nil {
		return generateErrorResponse(err.Error(), 500), nil
	}

	//Return result
	return events.APIGatewayProxyResponse{
		Body:       memberID,
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func generateErrorResponse(error string, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       error,
		StatusCode: code,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}

func main() {
	//Register handler
	lambda.Start(Handler)
}
