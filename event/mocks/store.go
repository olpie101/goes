// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_event is a generated GoMock package.
package mock_event

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	event "github.com/modernice/goes/event"
	time "github.com/modernice/goes/event/query/time"
	version "github.com/modernice/goes/event/query/version"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockStore) Insert(arg0 context.Context, arg1 ...event.Event) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockStoreMockRecorder) Insert(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockStore)(nil).Insert), varargs...)
}

// Find mocks base method
func (m *MockStore) Find(arg0 context.Context, arg1 uuid.UUID) (event.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(event.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockStore)(nil).Find), arg0, arg1)
}

// Query mocks base method
func (m *MockStore) Query(arg0 context.Context, arg1 event.Query) (event.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(event.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockStoreMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockStore)(nil).Query), arg0, arg1)
}

// Delete mocks base method
func (m *MockStore) Delete(arg0 context.Context, arg1 event.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), arg0, arg1)
}

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Names mocks base method
func (m *MockQuery) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names
func (mr *MockQueryMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockQuery)(nil).Names))
}

// IDs mocks base method
func (m *MockQuery) IDs() []uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDs")
	ret0, _ := ret[0].([]uuid.UUID)
	return ret0
}

// IDs indicates an expected call of IDs
func (mr *MockQueryMockRecorder) IDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDs", reflect.TypeOf((*MockQuery)(nil).IDs))
}

// Times mocks base method
func (m *MockQuery) Times() time.Constraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Times")
	ret0, _ := ret[0].(time.Constraints)
	return ret0
}

// Times indicates an expected call of Times
func (mr *MockQueryMockRecorder) Times() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Times", reflect.TypeOf((*MockQuery)(nil).Times))
}

// AggregateNames mocks base method
func (m *MockQuery) AggregateNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AggregateNames indicates an expected call of AggregateNames
func (mr *MockQueryMockRecorder) AggregateNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateNames", reflect.TypeOf((*MockQuery)(nil).AggregateNames))
}

// AggregateIDs mocks base method
func (m *MockQuery) AggregateIDs() []uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateIDs")
	ret0, _ := ret[0].([]uuid.UUID)
	return ret0
}

// AggregateIDs indicates an expected call of AggregateIDs
func (mr *MockQueryMockRecorder) AggregateIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateIDs", reflect.TypeOf((*MockQuery)(nil).AggregateIDs))
}

// AggregateVersions mocks base method
func (m *MockQuery) AggregateVersions() version.Constraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateVersions")
	ret0, _ := ret[0].(version.Constraints)
	return ret0
}

// AggregateVersions indicates an expected call of AggregateVersions
func (mr *MockQueryMockRecorder) AggregateVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateVersions", reflect.TypeOf((*MockQuery)(nil).AggregateVersions))
}

// Sortings mocks base method
func (m *MockQuery) Sortings() []event.SortOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sortings")
	ret0, _ := ret[0].([]event.SortOptions)
	return ret0
}

// Sortings indicates an expected call of Sortings
func (mr *MockQueryMockRecorder) Sortings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sortings", reflect.TypeOf((*MockQuery)(nil).Sortings))
}

// MockStream is a mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockStream) Next(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockStreamMockRecorder) Next(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStream)(nil).Next), arg0)
}

// Event mocks base method
func (m *MockStream) Event() event.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Event")
	ret0, _ := ret[0].(event.Event)
	return ret0
}

// Event indicates an expected call of Event
func (mr *MockStreamMockRecorder) Event() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockStream)(nil).Event))
}

// Err mocks base method
func (m *MockStream) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockStreamMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockStream)(nil).Err))
}

// Close mocks base method
func (m *MockStream) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStreamMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close), arg0)
}
