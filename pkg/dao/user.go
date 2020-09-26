package dao

import (
	"context"
	"database/sql"

	"github.com/rinchsan/gomock-todo/pkg/entity"
	"github.com/rinchsan/gomock-todo/pkg/repository"
)

func NewUser(db *sql.DB) repository.User {
	return userImpl{
		db: db,
	}
}

type userImpl struct {
	db *sql.DB
}

func (impl userImpl) Add(ctx context.Context, user *entity.User) error {
	panic("implement me")
}

func (impl userImpl) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	panic("implement me")
}

func (impl userImpl) GetAll(ctx context.Context) (entity.UserSlice, error) {
	panic("implement me")
}

func (impl userImpl) Update(ctx context.Context, user *entity.User) error {
	panic("implement me")
}
