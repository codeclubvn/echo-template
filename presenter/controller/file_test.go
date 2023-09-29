package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"trial_backend/pkg/lib"
	"trial_backend/usecase/mocks"
)

func TestFileController_SaveFile(t *testing.T) {
	cases := []struct {
		name    string
		req     *multipart.Form
		want    interface{}
		wantErr int
	}{
		{
			name: "happy flow",
			req: &multipart.Form{
				File: map[string][]*multipart.FileHeader{
					"file_request": {
						{
							Filename: "test.txt",
							Size:     100,
						},
					},
				},
			},
			wantErr: 200,
		},
	}

	// Setup
	e := echo.New()
	e.Validator = lib.NewRequestValidator()

	mockUCase := new(mocks.FileService)
	h := &FileController{
		fileService: mockUCase,
	}

	// Mocking
	mockUCase.On("SaveFile", mock.Anything, mock.Anything).Return(nil, nil)

	// Gọi hàm SaveFile của tầng handler
	for _, item := range cases {
		b, err := json.Marshal(item.req)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/v1/api/files", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		req.Header.Set("x-user-id", "123456")
		req.MultipartForm = item.req

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.SaveFile(c)) {
			assert.Equal(t, item.wantErr, rec.Code)
		}
	}

	// Kiểm tra rằng authService.Login đã được gọi một lần
	mockUCase.AssertCalled(t, "SaveFile", mock.Anything, mock.Anything)
}
