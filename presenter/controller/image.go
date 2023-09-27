package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"trail_backend/pkg/constants"
	utils2 "trail_backend/pkg/utils"
	"trail_backend/presenter/request"
	"trail_backend/usecase"
)

type FileCloudController struct {
	BaseController
	imageService usecase.FileCloudService
}

func NewImageController(imageService usecase.FileCloudService) *FileCloudController {
	return &FileCloudController{
		imageService: imageService,
	}
}

// Upload
//
//	@Summary		Upload
//	@Description	Upload
//	@Tags			Image
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			Authorization		header		string					true	"authorization token"
//	@Param			UploadFileRequest	formData	file					true	"UploadFileRequest"
//	@Success		200					{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/image/upload [POST]
func (h *FileCloudController) Upload(c echo.Context) error {
	var req request.UploadFileRequest
	if err := utils2.GetFile(c, &req, constants.FolderTmp); err != nil {
		return h.ResponseError(c, err)
	}
	defer os.Remove(req.FileName)

	data, err := h.imageService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	var res request.UploadFileResponse
	if err := utils2.Copy(&res, &data); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}
