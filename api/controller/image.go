package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"trail_backend/dto"
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

// Upload
// @Summary		Upload
// @Description	Upload
// @Tags		Image
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/image/upload [POST]
func (h *FileCloudController) Upload(c echo.Context) error {
	var req dto.UploadFileRequest
	if err := utils.GetFile(c, &req, constants.FolderTmp); err != nil {
		return h.ResponseError(c, err)
	}
	defer os.Remove(req.FileName)

	data, err := h.imageService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	var res dto.UploadFileResponse
	if err := utils.Copy(&res, &data); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}
