package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"trial_backend/pkg/constants"
	utils "trial_backend/pkg/utils"
	"trial_backend/presenter/request"
	"trial_backend/usecase"
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
//	@Security		Authorization
//
//	@Summary		SaveFile
//	@Description	SaveFile
//	@Tags			Image
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file_request	formData	file					true	"file_request"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/image/upload [POST]
func (h *FileCloudController) Upload(c echo.Context) error {
	file, err := c.FormFile("file_request")
	if err := c.Validate(file); err != nil {
		return h.ResponseValidatorError(c, err)
	}

	// check file is image
	if err := utils.CheckFileIsImage(file); err != nil {
		return h.ResponseError(c, err)
	}

	if _, err := os.Stat(constants.FolderTmp); os.IsNotExist(err) {
		if err := os.MkdirAll(constants.FolderTmp, 0755); err != nil {
			panic(err)
		}
	}
	if err := utils.GetFile(file, constants.FolderTmp); err != nil {
		return h.ResponseError(c, err)
	}
	defer os.Remove(constants.FolderTmp + file.Filename)

	data, err := h.imageService.Update(c.Request().Context(), file)
	if err != nil {
		return h.ResponseError(c, err)
	}

	var res request.UploadFileResponse
	if err := utils.Copy(&res, &data); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}
