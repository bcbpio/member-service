package repository

import (
	"errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"testing"
)

//Mock Neo4j Record
type mockRecord struct {
	MockRecord
}

func (m mockRecord) GetByIndex(index int) interface{} {
	return int64(0)
}

//Mock Neo4j Result
type mockResult struct {
	MockResult
}

func (m mockResult) Next() bool {
	return true
}

func (m mockResult) Record() neo4j.Record {
	return mockRecord{}
}

//Mock Neo4j Session
type mockSession struct {
	MockSession
}

func (mock mockSession) Run(cypher string, params map[string]interface{},
	configurers ...func(*neo4j.TransactionConfig)) (neo4j.Result, error) {

	if params["lastName"] == "" {
		return nil, errors.New("")
	}
	return mockResult{}, nil
}

func TestConnect(t *testing.T) {
	//Test Scenarios
	scenarios := []struct {
		//Global Parameters
		driverController  bool
		sessionController bool
		//Result
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	}{
		{true, true, MockDriver{}, MockSession{}, nil},
		{false, true, nil, nil, errors.New("")},
		{true, false, nil, nil, errors.New("")},
	}
	newDriver = mockNewDriver
	for index, scenario := range scenarios {
		mockDriverController = scenario.driverController
		mockSessionController = scenario.sessionController
		driver, session, err := Connect()
		if driver != scenario.driver ||
			session != scenario.session {
			t.Errorf("Test Case %d Failed - Expected '%v', '%v' Actual '%v', '%v'",
				index+1, scenario.driver, scenario.session, driver, session)
		}
		if (scenario.err == nil) && !errors.Is(err, scenario.err) {
			t.Errorf("Test Case %d Error Failed - Expected '%v' | Actual '%v'",
				index+1, scenario.err, err)
		} else if (scenario.err != nil) && !errors.As(err, &scenario.err) {
			t.Errorf("Test Case %d Error Failed - Expected '%T' | Actual '%T'",
				index+1, scenario.err, err)
		}
	}
}

func TestCreateMember(t *testing.T) {
	scenarios := []struct {
		//Parameters
		m Member
		//Result
		result string
		err    error
	}{
		{Member{}, "", errors.New("")},
		{Member{LastName: "MOCK_LNAME"}, "0", nil},
	}
	repo := NewRepository(mockSession{})
	for index, scenario := range scenarios {
		id, err := repo.CreateMember(scenario.m)
		if id != scenario.result {
			t.Errorf("Test Case %d Failed - Expected '%s' | Actual '%s'",
				index+1, scenario.result, id)
		}
		if (scenario.err == nil) && !errors.Is(err, scenario.err) {
			t.Errorf("Test Case %d Error Failed - Expected '%v' | Actual '%v'",
				index+1, scenario.err, err)
		} else if (scenario.err != nil) && !errors.As(err, &scenario.err) {
			t.Errorf("Test Case %d Error Failed - Expected '%T' | Actual '%T'",
				index+1, scenario.err, err)
		}
	}
}
