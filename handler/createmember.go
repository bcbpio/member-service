package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bcbpio/member-service/repository"
	"github.com/bcbpio/member-service/service"
)

//DB connection function
var Connect = repository.Connect

//New Repository Initializer
var NewRepository = repository.NewRepository

//New Service Initializer
var NewService = service.NewService

//Handler - handler for create member
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Connect to db
	driver, session, err := Connect()
	if err != nil {
		return generateErrorResponse(err.Error(), 500), err
	}
	defer driver.Close()
	defer session.Close()

	//Get new contact service
	var memberSvc service.Service
	{
		repo := NewRepository(session)
		memberSvc = NewService(repo)
	}

	//Create member using post request body as query parameter
	memberID, err := memberSvc.CreateMember(req.Body)
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

func Main() {
	//Register handler
	lambda.Start(Handler)
}
