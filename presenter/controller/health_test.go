package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthController_Health(t *testing.T) {
	t.Helper()

	// Tạo một đối tượng Echo để sử dụng trong unit test
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/heath", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Tạo một đối tượng HealthController
	h := &HealthController{}

	// Gọi hàm Health
	if assert.NoError(t, h.Health(c)) {
		// Kiểm tra mã trạng thái HTTP
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
