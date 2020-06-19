package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bcbpio/member-service/repository"
	"github.com/bcbpio/member-service/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"os"
)

//Handler - handler for create member
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Get database connection
	//DB Variables
	dbHost := os.Getenv("db_host")
	dbUsername := os.Getenv("db_username")
	dbPassword := os.Getenv("db_password")
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
	driver, err := neo4j.NewDriver(dbHost, neo4j.BasicAuth(dbUsername, dbPassword, ""), configForNeo4j40)
	if err != nil {
		return generateErrorResponse(err.Error(), 500), err
	}
	defer driver.Close()

	//Get new contact service
	var memberSvc service.Service
	{
		repo := repository.NewRepository(driver)
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
