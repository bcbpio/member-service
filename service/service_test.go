package service

import (
	"fmt"
	"github.com/bcbpio/member-service/repository"
	"testing"
)

type mockError struct {
	mockMessage string
}

func (mockE mockError) Error() string {
	return mockE.mockMessage
}

//Mock Repository
type mockRepository struct{}

//Mock Create Member Function
func (mock *mockRepository) CreateMember(m repository.Member) (string, error) {
	if m.LastName == "" {
		return "", mockError{}
	}
	return "0", nil
}

func TestCreateMember(t *testing.T) {
	//Test Scenarios
	scenarios := []struct {
		//Parameters
		m repository.Member
		//Result
		result string
		err    error
	}{
		{repository.Member{LastName: "MockLastName"}, "0", nil},
		{repository.Member{}, "", mockError{}},
	}

	//Run Test Cases
	svc := NewService(&mockRepository{})
	for index, scenario := range scenarios {
		result, err := svc.CreateMember(scenario.m)
		//Result Check
		fmt.Print(result, err)
		if result != scenario.result || err != scenario.err {
			t.Errorf("Test Case %d Failed - Expected %s, %v Actual %s, %v",
				index, scenario.result, scenario.err, result, err)
		}
	}
}
