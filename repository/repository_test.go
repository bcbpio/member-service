package repository

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"testing"
)

//Mock Neo4j Record
type mockRecord struct{}

func (m mockRecord) Keys() []string {
	panic("implement me")
}

func (m mockRecord) Values() []interface{} {
	panic("implement me")
}

func (m mockRecord) Get(key string) (interface{}, bool) {
	panic("implement me")
}

func (m mockRecord) GetByIndex(index int) interface{} {
	return int64(0)
}

//Mock Neo4j Result
type mockResult struct{}

func (m mockResult) Keys() ([]string, error) {
	panic("implement me")
}

func (m mockResult) Next() bool {
	return true
}

func (m mockResult) Err() error {
	panic("implement me")
}

func (m mockResult) Record() neo4j.Record {
	return mockRecord{}
}

func (m mockResult) Summary() (neo4j.ResultSummary, error) {
	panic("implement me")
}

func (m mockResult) Consume() (neo4j.ResultSummary, error) {
	panic("implement me")
}

//Mock Neo4j Session
type mockSession struct{}

func (mock mockSession) LastBookmark() string {
	panic("implement me")
}

func (mock mockSession) BeginTransaction(configurers ...func(*neo4j.TransactionConfig)) (neo4j.Transaction, error) {
	panic("implement me")
}

func (mock mockSession) ReadTransaction(work neo4j.TransactionWork, configurers ...func(*neo4j.TransactionConfig)) (interface{}, error) {
	panic("implement me")
}

func (mock mockSession) WriteTransaction(work neo4j.TransactionWork, configurers ...func(*neo4j.TransactionConfig)) (interface{}, error) {
	panic("implement me")
}

func (mock mockSession) Close() error {
	panic("implement me")
}

func (mock mockSession) Run(cypher string, params map[string]interface{},
	configurers ...func(*neo4j.TransactionConfig)) (neo4j.Result, error) {

	if params["lastName"] == "" {
		return nil, mockError{"Incomplete Payload"}
	}
	return mockResult{}, nil
}

type mockError struct {
	mockMessage string
}

func (mockE mockError) Error() string {
	return mockE.mockMessage
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
