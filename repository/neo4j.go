package repository

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"os"
	"strconv"
)

type repository struct {
	db neo4j.Session
}

//NewRepository - instantiate default service repository
func NewRepository(db neo4j.Session) Repository {
	return &repository{
		db: db,
	}
}

var newDriver = neo4j.NewDriver

func Connect() (neo4j.Driver, neo4j.Session, error) {
	//Get database connection
	//DB Variables
	dbHost := os.Getenv("db_host")
	dbUsername := os.Getenv("db_username")
	dbPassword := os.Getenv("db_password")
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
	driver, err := newDriver(dbHost, neo4j.BasicAuth(dbUsername, dbPassword, ""), configForNeo4j40)
	if err != nil {
		return nil, nil, err
	}
	//Start session
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return driver, nil, err
	}

	return driver, session, nil
}

func (r *repository) CreateMember(m Member) (string, error) {
	cypherStatement := `
		CREATE (m:Member{lastName: $lastName, firstName: $firstName,
			nickname: $nickname, homeNo: $homeNo, businessNo: $businessNo,
			mobileNo: $mobileNo, birthdate: $birthdate, residence: $residence,
			emailAddress: $emailAddress})
		RETURN ID(m)
	`

	result, err := r.db.Run(cypherStatement, map[string]interface{}{
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
