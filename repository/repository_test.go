package repository

import (
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
		return nil, mockError{"Incomplete Payload"}
	}
	return mockResult{}, nil
}

func TestConnect(t *testing.T) {

	newDriver = mockNewDriver
	driver, session, err := Connect()

	if driver == nil {
		t.Error("Should not be nil")
	}

	if session == nil {
		t.Error("Should not be nil")
	}

	if err != nil {
		t.Error("Should not be error")
	}

	mockDriverController = false
	driver, _, err = Connect()
	if driver != nil {
		t.Error("Should be nil")
	}
	if err == nil {
		t.Error("Should be error")
	}

	mockDriverController = true
	mockSessionController = false
	_, session, _ = Connect()
	if session != nil {
		t.Error("Should be nil")
	}
}

func TestCreateMember(t *testing.T) {
	repo := NewRepository(&mockSession{})
	id, err := repo.CreateMember(Member{})

	if id != "" {
		t.Error("Wrong id result")
	}
	if err == nil {
		t.Error("Should be error")
	}

	id, err = repo.CreateMember(Member{LastName: "MOCK_LNAME"})

	if id != "0" {
		t.Error("Wrong id result")
	}
	if err != nil {
		t.Error("Should not be error")
	}

}
