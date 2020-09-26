package registry

import (
	"github.com/golang/mock/gomock"
	"github.com/rinchsan/gomock-todo/pkg/repository"
	"github.com/rinchsan/gomock-todo/pkg/repository/mock"
)

func NewMockRepository(ctrl *gomock.Controller) Repository {
	return mockRepositoryImpl{
		ctrl: ctrl,
	}
}

type mockRepositoryImpl struct {
	ctrl *gomock.Controller
}

func (impl mockRepositoryImpl) User() repository.User {
	return mock.NewUser(impl.ctrl)
}

func (impl mockRepositoryImpl) Todo() repository.Todo {
	return mock.NewTodo(impl.ctrl)
}
