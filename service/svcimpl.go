package service

import (
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

func (s *service) CreateMember(m repository.Member) (string, error) {

	//Create contact via injected repository
	fmt.Println(m.LastName)
	id, err := s.repository.CreateMember(m)
	if err != nil {
		return "", err
	}

	return id, nil
}
