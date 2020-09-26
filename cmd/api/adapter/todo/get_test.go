package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/rinchsan/gomock-todo/pkg/entity"
	"github.com/rinchsan/gomock-todo/pkg/registry"
	"github.com/rinchsan/gomock-todo/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAll(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) handler
		code  int
	}{
		"repository.Todo.GetAll returns error": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				h.todo.(*mock.Todo).EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("test error"))
				return h
			},
			code: http.StatusInternalServerError,
		},
		"repository.Todo.GetAll succeeds": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				h.todo.(*mock.Todo).EXPECT().GetAll(gomock.Any()).Return(entity.TodoSlice{}, nil)
				return h
			},
			code: http.StatusOK,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := c.setup(ctrl)

			r := httptest.NewRequest(http.MethodGet, "/todos", nil)
			w := httptest.NewRecorder()
			h.GetAll(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
