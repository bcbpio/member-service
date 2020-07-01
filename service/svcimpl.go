package service

import (
	"encoding/json"
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
		return "", err
	}

	//Create contact via injected repository
	id, err := s.repository.CreateMember(m)
	if err != nil {
		return "", err
	}

	return id, nil
}
