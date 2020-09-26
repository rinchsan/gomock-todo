package dao

import (
	"context"
	"database/sql"

	"github.com/rinchsan/gomock-todo/pkg/entity"
	"github.com/rinchsan/gomock-todo/pkg/repository"
)

func NewTodo(db *sql.DB) repository.Todo {
	return todoImpl{
		db: db,
	}
}

type todoImpl struct {
	db *sql.DB
}

func (impl todoImpl) Add(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error {
	panic("implement me")
}

func (impl todoImpl) GetByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	panic("implement me")
}

func (impl todoImpl) GetAll(ctx context.Context) (entity.TodoSlice, error) {
	panic("implement me")
}

func (impl todoImpl) Update(ctx context.Context, todo *entity.Todo, assigneeUserIDs []uint64) error {
	panic("implement me")
}
