package controller

import (
	"echo_template/pkg/lib"
	"echo_template/usecase/mocks"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileController_SaveFile(t *testing.T) {
	cases := []struct {
		name     string
		nameFile string
		want     interface{}
		wantErr  int
	}{
		{
			name:     "happy flow",
			nameFile: "test.png",
			wantErr:  201,
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
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		go func() {
			defer writer.Close()
			part, err := writer.CreateFormFile("file_request", item.name)
			if err != nil {
				t.Error(err)
			}
			img := createImage()
			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()

		req := httptest.NewRequest(http.MethodPost, "/v1/api/files", pr)
		req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		req.Header.Set("x-user-id", uuid.NewV4().String())

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.SaveFile(c)) {
			assert.Equal(t, item.wantErr, rec.Code)
		}
	}

	// Kiểm tra rằng authService.Login đã được gọi một lần
	mockUCase.AssertCalled(t, "SaveFile", mock.Anything, mock.Anything)
}

func createImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	return img
}
