package service

//Service - interface for required business logic to be implemented
type Service interface {
	CreateMember(requestBody string) (string, error)
}
