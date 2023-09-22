package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"trail_backend/api/dto"
	"trail_backend/service"
	"trail_backend/utils"
	"trail_backend/utils/constants"
)

type FileCloudController struct {
	BaseController
	imageService service.FileCloudService
}

func NewImageController(imageService service.FileCloudService) *FileCloudController {
	return &FileCloudController{
		imageService: imageService,
	}
}

func (h *FileCloudController) Upload(c echo.Context) error {
	var req dto.UploadFileRequest
	if err := utils.GetFile(c, &req, constants.FolderTmp); err != nil {
		return h.ResponseError(c, err)
	}
	_ = os.Remove(req.FileName)

	_, err := h.imageService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", nil)
}
