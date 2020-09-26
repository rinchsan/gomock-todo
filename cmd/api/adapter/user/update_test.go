package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
		"repository.User.GetByID returns error": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				h.user.(*mock.User).EXPECT().GetByID(gomock.Any(), uint64(1)).Return(nil, errors.New("test error"))
				return h
			},
			body: `{"id":1, "username":"rinchsan"}`,
			code: http.StatusInternalServerError,
		},
		"repository.User.Update returns error": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				user := &entity.User{ID: 1}
				h.user.(*mock.User).EXPECT().GetByID(gomock.Any(), uint64(1)).Return(user, nil)
				user.Username = "rinchsan"
				h.user.(*mock.User).EXPECT().Update(gomock.Any(), user).Return(errors.New("test error"))
				return h
			},
			body: `{"id":1, "username":"rinchsan"}`,
			code: http.StatusInternalServerError,
		},
		"repository.User.Update succeeds": {
			setup: func(ctrl *gomock.Controller) handler {
				h := newHandler(registry.NewMockRepository(ctrl))
				user := &entity.User{ID: 1}
				h.user.(*mock.User).EXPECT().GetByID(gomock.Any(), uint64(1)).Return(user, nil)
				user.Username = "rinchsan"
				h.user.(*mock.User).EXPECT().Update(gomock.Any(), user).Return(nil)
				return h
			},
			body: `{"id":1, "username":"rinchsan"}`,
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

			r := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(c.body))
			w := httptest.NewRecorder()
			h.Update(w, r)
			assert.Equal(t, c.code, w.Code)
		})
	}
}
