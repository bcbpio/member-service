package repository

import (
	"time"
)

type Member struct {
	ID           int       `json:"id"`
	LastName     string    `json:"lastName"`
	FirstName    string    `json:"firstName"`
	NickName     string    `json:"nickname"`
	HomeNo       string    `json:"homeNo"`
	BusinessNo   string    `json:"businessNo"`
	MobileNo     string    `json:"mobileNo"`
	Birthdate    time.Time `json:"birthdate"`
	Residence    string    `json:"residence"`
	EmailAddress string    `json:"emailAddress"`
}

//Repository - interface for required methods to be implemented
type Repository interface {
	CreateMember(m Member) (string, error)
}
