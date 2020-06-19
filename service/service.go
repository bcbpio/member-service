package service

import "github.com/bcbpio/member-service/repository"

//Service - interface for required business logic to be implemented
type Service interface {
	CreateMember(m repository.Member) (string, error)
}
