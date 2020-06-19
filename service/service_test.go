package service

import (
	"github.com/bcbpio/member-service/repository"
	"testing"
)

type mockError struct {
	mockMessage string
}

func (mockE mockError) Error() string {

	return mockE.mockMessage
}

type mockRepository struct{}

func (mock *mockRepository) CreateMember(m repository.Member) (string, error) {
	if m.LastName == "" {
		return "", mockError{"Incomplete Payload"}
	}
	return "0", nil
}

func TestCreateMember(t *testing.T) {
	svc := NewService(&mockRepository{})
	//Test success
	result, err := svc.CreateMember(repository.Member{LastName: "Lim"})

	if result != "0" {
		t.Error("Wrong service result")
	}
	if err != nil {
		t.Error("Should not be error")
	}

	//Test failure
	result, err = svc.CreateMember(repository.Member{})
	if err == nil {
		t.Error("Should be error")
	}
	if err != nil && err.Error() != "Incomplete Payload" {
		t.Error("Wrong error message")
	}
	if result == "" {
		t.Error("Wrong service result")
	}
}
