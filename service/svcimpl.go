package service

import (
	"encoding/json"
	"fmt"
	"github.com/bcbpio/member-service/repository"
)

type service struct {
	repository repository.Repository
}

//NewService - instantiate service
func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateMember(requestBody string) (string, error) {
	//Unmarshal request string
	var m repository.Member
	err := json.Unmarshal([]byte(requestBody), &m)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//Create contact via injected repository
	fmt.Println(m.LastName)
	id, err := s.repository.CreateMember(m)
	if err != nil {
		return "", err
	}

	return id, nil
}
