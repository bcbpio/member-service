package service

import (
	"errors"
	"github.com/bcbpio/member-service/repository"
	"testing"
)

//Mock Repository
type mockRepository struct{}

//Mock Create Member Function
func (mock *mockRepository) CreateMember(m repository.Member) (string, error) {
	if m.LastName == "" {
		return "", errors.New("")
	}
	return "0", nil
}

func TestCreateMember(t *testing.T) {
	//Test Scenarios
	scenarios := []struct {
		//Parameters
		requestBody string
		//Result
		result string
		err    error
	}{
		{"{\"lastName\":\"MOCK\"}", "0", nil},
		{"{\"lastName\":\"\"}", "", errors.New("")},
		{"", "", errors.New("unexpected end of JSON input")},
	}

	//Run Test Cases
	svc := NewService(&mockRepository{})
	for index, scenario := range scenarios {
		result, err := svc.CreateMember(scenario.requestBody)
		if result != scenario.result && !errors.Is(err, scenario.err) {
			t.Errorf("Test Case %d Failed - Expected '%s', '%v' Actual '%s', '%v'",
				index+1, scenario.result, scenario.err, result, err)
		}
	}
}
