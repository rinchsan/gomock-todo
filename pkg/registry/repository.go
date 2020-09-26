package registry

import (
	"github.com/rinchsan/gomock-todo/pkg/conn"
	"github.com/rinchsan/gomock-todo/pkg/dao"
	"github.com/rinchsan/gomock-todo/pkg/repository"
)

func NewRepository() (Repository, func()) {
	closeDB := conn.SetupDB()
	return repositoryImpl{}, func() {
		closeDB()
	}
}

type Repository interface {
	User() repository.User
	Todo() repository.Todo
}

type repositoryImpl struct {
}

func (impl repositoryImpl) User() repository.User {
	return dao.NewUser(conn.DB)
}

func (impl repositoryImpl) Todo() repository.Todo {
	return dao.NewTodo(conn.DB)
}
