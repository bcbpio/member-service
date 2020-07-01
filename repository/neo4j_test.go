package repository

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/url"
)

var mockDriverController = true

//Mock NewDriver Method
var mockNewDriver = func(target string, auth neo4j.AuthToken, configurers ...func(*neo4j.Config)) (neo4j.Driver, error) {
	for _, configurer := range configurers {
		configurer(&neo4j.Config{})
	}
	if mockDriverController {
		return MockDriver{}, nil
	}
	return nil, mockError{}
}

//Mock Neo4j Driver
type MockDriver struct{}

func (mockDriver MockDriver) Target() url.URL {
	return url.URL{}
}

func (mockDriver MockDriver) Session(accessMode neo4j.AccessMode, bookmarks ...string) (neo4j.Session, error) {
	return nil, nil
}

var mockSessionController = true

func (mockDriver MockDriver) NewSession(config neo4j.SessionConfig) (neo4j.Session, error) {
	if mockSessionController {
		return MockSession{}, nil
	}
	return nil, mockError{}
}

func (mockDriver MockDriver) Close() error {
	return nil
}

//Mock Neo4j Record
type MockRecord struct{}

func (mockRecord MockRecord) Keys() []string {
	return nil
}

func (mockRecord MockRecord) Values() []interface{} {
	return nil
}

func (mockRecord MockRecord) Get(key string) (interface{}, bool) {
	return nil, true
}

func (mockRecord MockRecord) GetByIndex(index int) interface{} {
	return nil
}

//Mock Neo4j Result
type MockResult struct{}

func (mockResult MockResult) Keys() ([]string, error) {
	return nil, nil
}

func (mockResult MockResult) Next() bool {
	return true
}

func (mockResult MockResult) Err() error {
	return nil
}

func (mockResult MockResult) Record() neo4j.Record {
	return nil
}

func (mockResult MockResult) Summary() (neo4j.ResultSummary, error) {
	return nil, nil
}

func (mockResult MockResult) Consume() (neo4j.ResultSummary, error) {
	return nil, nil
}

//Mock Neo4j Session
type MockSession struct{}

func (mockSession MockSession) LastBookmark() string {
	return ""
}

func (mockSession MockSession) BeginTransaction(configurers ...func(*neo4j.TransactionConfig)) (neo4j.Transaction, error) {
	return nil, nil
}

func (mockSession MockSession) ReadTransaction(work neo4j.TransactionWork, configurers ...func(*neo4j.TransactionConfig)) (interface{}, error) {
	return nil, nil
}

func (mockSession MockSession) WriteTransaction(work neo4j.TransactionWork, configurers ...func(*neo4j.TransactionConfig)) (interface{}, error) {
	return nil, nil
}

func (mockSession MockSession) Close() error {
	return nil
}

func (mockSession MockSession) Run(cypher string, params map[string]interface{},
	configurers ...func(*neo4j.TransactionConfig)) (neo4j.Result, error) {
	return nil, nil
}

//Mock error
type mockError struct {
	mockMessage string
}

func (mockE mockError) Error() string {
	return mockE.mockMessage
}
