package handler

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"os"
)

func Connect() (neo4j.Driver, neo4j.Session, error) {
	//Get database connection
	//DB Variables
	dbHost := os.Getenv("db_host")
	dbUsername := os.Getenv("db_username")
	dbPassword := os.Getenv("db_password")
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
	driver, err := neo4j.NewDriver(dbHost, neo4j.BasicAuth(dbUsername, dbPassword, ""), configForNeo4j40)
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
