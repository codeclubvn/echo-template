package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"trial_backend/pkg/lib"
	"trial_backend/presenter/request"
	"trial_backend/usecase/mocks"
)

func TestPostController_Create(t *testing.T) {
	t.Helper()

	cases := []struct {
		name    string
		req     request.CreatePostRequest
		want    interface{}
		wantErr int
	}{
		{
			name: "happy flow",
			req: request.CreatePostRequest{
				Title:   "New Title",
				Content: "New Content",
				Slug:    "new-slug",
				Image:   "new-image",
				Files:   []string{uuid.NewV4().String(), uuid.NewV4().String()},
			},
			wantErr: 201,
		},
		{
			name: "invalid Files",
			req: request.CreatePostRequest{
				Title:   "New Title",
				Content: "New Content",
				Slug:    "new-slug",
				Image:   "new-image",
				// Files không hợp lệ
				Files: []string{"invalid-file"},
			},
			wantErr: 422,
		},
	}

	// Tạo một đối tượng Echo để sử dụng trong unit test
	e := echo.New()
	e.Validator = lib.NewRequestValidator()

	// Tạo một đối tượng PostController
	mockUCase := new(mocks.PostService)
	h := &PostController{
		postService: mockUCase, // mockPostService là một mock object cho PostService
	}

	// Mocking
	mockUCase.On("Create", mock.Anything, mock.Anything).Return(nil, nil)

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			b, err := json.Marshal(item.req)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/api/posts", strings.NewReader(string(b)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set("x-user-id", uuid.NewV4().String())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, h.Create(c)) {
				assert.Equal(t, item.wantErr, rec.Code)
			}
		})
		// Kiểm tra rằng authService.Login đã được gọi một lần
		mockUCase.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	}
}

func TestPostController_Update(t *testing.T) {
	t.Helper()

	cases := []struct {
		name    string
		req     request.UpdatePostRequest
		want    interface{}
		wantErr int
	}{
		{
			name: "happy flow",
			req: request.UpdatePostRequest{
				ID: uuid.NewV4().String(),
				CreatePostRequest: request.CreatePostRequest{
					Title:   "New Title",
					Content: "New Content",
					Slug:    "new-slug",
					Image:   "new-image",
					Files:   []string{uuid.NewV4().String(), uuid.NewV4().String()},
				},
			},
			wantErr: 200,
		},
		{
			name: "invalid ID",
			req: request.UpdatePostRequest{
				// ID thiếu
				CreatePostRequest: request.CreatePostRequest{
					Title:   "New Title",
					Content: "New Content",
					Slug:    "new-slug",
					Image:   "new-image",
					Files:   []string{uuid.NewV4().String(), uuid.NewV4().String()},
				},
			},
			wantErr: 422,
		},
		{
			name: "invalid Files",
			req: request.UpdatePostRequest{
				ID: "valid-post-id",
				CreatePostRequest: request.CreatePostRequest{
					Title:   "New Title",
					Content: "New Content",
					Slug:    "new-slug",
					Image:   "new-image",
					// Files không hợp lệ
					Files: []string{"invalid-file"},
				},
			},
			wantErr: 422,
		},
	}

	// Tạo một đối tượng Echo để sử dụng trong unit test
	e := echo.New()
	e.Validator = lib.NewRequestValidator()

	// Tạo một đối tượng PostController
	mockUCase := new(mocks.PostService)
	h := &PostController{
		postService: mockUCase, // mockPostService là một mock object cho PostService
	}

	// Mocking
	mockUCase.On("Update", mock.Anything, mock.Anything).Return(nil, nil)

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			b, err := json.Marshal(item.req)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPut, "/v1/api/posts", strings.NewReader(string(b)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set("x-user-id", uuid.NewV4().String())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, h.Update(c)) {
				assert.Equal(t, item.wantErr, rec.Code)
			}
		})
		// Kiểm tra rằng authService.Login đã được gọi một lần
		mockUCase.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	}
}
