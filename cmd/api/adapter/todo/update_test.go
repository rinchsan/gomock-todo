package todo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/rinchsan/gomock-todo/pkg/entity"
	"github.com/rinchsan/gomock-todo/pkg/registry"
	"github.com/rinchsan/gomock-todo/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Update(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) handler
		body  string
		code  int
	}{
		"invalid json body": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				return h
			},
			body: `{{}`,
			code: http.StatusBadRequest,
		},
		"repository.Todo.GetByID returns error": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				h.todo.(*mock.Todo).EXPECT().GetByID(gomock.Any(), uint64(1)).Return(nil, errors.New("test error"))
				return h
			},
			body: `{"id":1, "title":"new title", "detail":"new detail", "due_date":"2020-08-31T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}`,
			code: http.StatusInternalServerError,
		},
		"repository.Todo.Update returns error": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				todo := &entity.Todo{ID: 1}
				h.todo.(*mock.Todo).EXPECT().GetByID(gomock.Any(), uint64(1)).Return(todo, nil)
				todo.Title = "todo title"
				todo.Detail = "todo detail"
				todo.AuthorUserID = 3
				todo.DueDate = time.Date(2020, time.August, 31, 0, 0, 0, 0, time.UTC)
				h.todo.(*mock.Todo).EXPECT().Update(gomock.Any(), todo, []uint64{2, 3}).Return(errors.New("test error"))
				return h
			},
			body: `{"id":1, "title":"new title", "detail":"new detail", "due_date":"2020-08-31T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}`,
			code: http.StatusInternalServerError,
		},
		"repository.Todo.Update succeeds": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				todo := &entity.Todo{ID: 1}
				h.todo.(*mock.Todo).EXPECT().GetByID(gomock.Any(), uint64(1)).Return(todo, nil)
				todo.Title = "todo title"
				todo.Detail = "todo detail"
				todo.AuthorUserID = 3
				todo.DueDate = time.Date(2020, time.August, 31, 0, 0, 0, 0, time.UTC)
				h.todo.(*mock.Todo).EXPECT().Update(gomock.Any(), todo, []uint64{2, 3}).Return(nil)
				return h
			},
			body: `{"id":1, "title":"new title", "detail":"new detail", "due_date":"2020-08-31T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}`,
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

			r := httptest.NewRequest(http.MethodPut, "/todos", strings.NewReader(c.body))
			w := httptest.NewRecorder()
			h.Update(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
