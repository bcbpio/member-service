package handler

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/bcbpio/member-service/repository"
	"github.com/bcbpio/member-service/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/url"
	"reflect"
	"testing"
)

type mockService struct{}

func (m *mockService) CreateMember(reqBody string) (string, error) {
	if reqBody == "" {
		return "", errors.New("")
	}
	return "0", nil
}

//Mock Repository
type mockRepository struct{}

//Mock Create Member Function
func (mock *mockRepository) CreateMember(m repository.Member) (string, error) {
	return "", nil
}

type mockDriver struct{}

func (m mockDriver) Target() url.URL {
	panic("implement me")
}

func (m mockDriver) Session(accessMode neo4j.AccessMode, bookmarks ...string) (neo4j.Session, error) {
	panic("implement me")
}

func (m mockDriver) NewSession(config neo4j.SessionConfig) (neo4j.Session, error) {
	panic("implement me")
}

func (m mockDriver) Close() error {
	return nil
}

type mockSession struct{}

func (m mockSession) LastBookmark() string {
	panic("implement me")
}

func (m mockSession) BeginTransaction(configurers ...func(*neo4j.TransactionConfig)) (neo4j.Transaction, error) {
	panic("implement me")
}

func (m mockSession) ReadTransaction(work neo4j.TransactionWork, configurers ...func(*neo4j.TransactionConfig)) (interface{}, error) {
	panic("implement me")
}

func (m mockSession) WriteTransaction(work neo4j.TransactionWork, configurers ...func(*neo4j.TransactionConfig)) (interface{}, error) {
	panic("implement me")
}

func (m mockSession) Run(cypher string, params map[string]interface{}, configurers ...func(*neo4j.TransactionConfig)) (neo4j.Result, error) {
	panic("implement me")
}

func (m mockSession) Close() error {
	return nil
}

func TestCreateMemberHandler(t *testing.T) {
	forceError := false
	Connect = func() (neo4j.Driver, neo4j.Session, error) {
		if forceError {
			return nil, nil, errors.New("")
		}
		return mockDriver{}, mockSession{}, nil
	}
	NewRepository = func(db neo4j.Session) repository.Repository {
		return &mockRepository{}
	}
	NewService = func(repository repository.Repository) service.Service {
		return &mockService{}
	}

	scenarios := []struct {
		//Global Parameter
		forceError bool
		//Parameter
		request events.APIGatewayProxyRequest
		//Return
		response events.APIGatewayProxyResponse
		err      error
	}{
		{
			false,
			events.APIGatewayProxyRequest{Body: ""},
			events.APIGatewayProxyResponse{
				Body:       "",
				StatusCode: 500,
				Headers: map[string]string{
					"Access-Control-Allow-Origin":      "*",
					"Access-Control-Allow-Credentials": "true",
				},
			},
			nil,
		},
		{
			false,
			events.APIGatewayProxyRequest{Body: "{lastName:''}"},
			events.APIGatewayProxyResponse{
				Body:       "0",
				StatusCode: 200,
				Headers: map[string]string{
					"Access-Control-Allow-Origin":      "*",
					"Access-Control-Allow-Credentials": "true",
				},
			},
			nil,
		},
		{
			true,
			events.APIGatewayProxyRequest{},
			events.APIGatewayProxyResponse{
				Body:       "",
				StatusCode: 500,
				Headers: map[string]string{
					"Access-Control-Allow-Origin":      "*",
					"Access-Control-Allow-Credentials": "true",
				},
			},
			nil,
		},
	}
	for index, scenario := range scenarios {
		forceError = scenario.forceError
		response, _ := CreateMemberHandler(scenario.request)
		if !reflect.DeepEqual(response, scenario.response) {
			t.Errorf("Test Case %d Result Failed - Expected '%v' | Actual '%v'",
				index+1, scenario.response, response)
		}
	}
}
