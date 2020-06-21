package repository

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"strconv"
)

type repository struct {
	db neo4j.Driver
}

//NewRepository - instantiate default service repository
func NewRepository(db interface{}) Repository {
	return &repository{
		db: db.(neo4j.Driver),
	}
}

func (r *repository) CreateMember(m Member) (string, error) {
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := r.db.NewSession(sessionConfig)
	if err != nil {
		return "", err
	}
	defer session.Close()

	cypherStatement := `
		CREATE (m:Member{lastName: $lastName, firstName: $firstName,
			nickname: $nickname, homeNo: $homeNo, businessNo: $businessNo,
			mobileNo: $mobileNo, birthdate: $birthdate, residence: $residence,
			emailAddress: $emailAddress})
		RETURN ID(m)
	`

	result, err := session.Run(cypherStatement, map[string]interface{}{
		"lastName":     m.LastName,
		"firstName":    m.FirstName,
		"nickname":     m.NickName,
		"homeNo":       m.HomeNo,
		"businessNo":   m.BusinessNo,
		"mobileNo":     m.MobileNo,
		"birthdate":    m.Birthdate,
		"residence":    m.Residence,
		"emailAddress": m.EmailAddress,
	})
	if err != nil {
		return "", err
	}

	result.Next()
	memberID := result.Record().GetByIndex(0).(int64)

	return strconv.Itoa(int(memberID)), nil
}
