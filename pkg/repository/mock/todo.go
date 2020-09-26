// Code generated by MockGen. DO NOT EDIT.
// Source: ./todo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/rinchsan/gomock-todo/pkg/entity"
	reflect "reflect"
)

// Todo is a mock of Todo interface
type Todo struct {
	ctrl     *gomock.Controller
	recorder *TodoMockRecorder
}

// TodoMockRecorder is the mock recorder for Todo
type TodoMockRecorder struct {
	mock *Todo
}

// NewTodo creates a new mock instance
func NewTodo(ctrl *gomock.Controller) *Todo {
	mock := &Todo{ctrl: ctrl}
	mock.recorder = &TodoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Todo) EXPECT() *TodoMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *Todo) Add(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, todo, assigneeUserIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *TodoMockRecorder) Add(ctx, todo, assigneeUserIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*Todo)(nil).Add), ctx, todo, assigneeUserIDs)
}

// GetByID mocks base method
func (m *Todo) GetByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *TodoMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*Todo)(nil).GetByID), ctx, id)
}

// GetAll mocks base method
func (m *Todo) GetAll(ctx context.Context) (entity.TodoSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(entity.TodoSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *TodoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*Todo)(nil).GetAll), ctx)
}

// Update mocks base method
func (m *Todo) Update(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, todo, assigneeUserIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *TodoMockRecorder) Update(ctx, todo, assigneeUserIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Todo)(nil).Update), ctx, todo, assigneeUserIDs)
}
