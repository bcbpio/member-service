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
		{"", "", errors.New("")},
	}

	//Run Test Cases
	svc := NewService(&mockRepository{})
	for index, scenario := range scenarios {
		result, err := svc.CreateMember(scenario.requestBody)
		if result != scenario.result {
			t.Errorf("Test Case %d Result Failed - Expected '%s' | Actual '%s'",
				index+1, scenario.result, result)
		}
		if (scenario.err == nil) && !errors.Is(err, scenario.err) {
			t.Errorf("Test Case %d Error Failed - Expected '%v' | Actual '%v'",
				index+1, scenario.err, err)
		} else if (scenario.err != nil) && !errors.As(err, &scenario.err) {
			t.Errorf("Test Case %d Error Failed - Expected '%T' | Actual '%T'",
				index+1, scenario.err, err)
		}
	}
}
