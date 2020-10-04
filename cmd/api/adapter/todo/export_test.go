package todo

import "github.com/rinchsan/gomock-todo/pkg/repository/mock"

type Handler = handler

func NewHandler(
	user *mock.User,
	todo *mock.Todo,
) Handler {
	return Handler{
		user: user,
		todo: todo,
	}
}
