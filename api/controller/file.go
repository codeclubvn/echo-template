package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"trail_backend/api/dto"
	"trail_backend/service"
	"trail_backend/utils"
)

type FileController struct {
	BaseController
	fileService service.FileService
}

func NewFileController(fileService service.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

// Upload
// @Summary		Upload
// @Description	Upload
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/file [Post]
func (h *FileController) Upload(c echo.Context) error {
	var req dto.UploadFileRequest
	file, err := c.FormFile("file")
	req.File = file

	data, err := h.fileService.Upload(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", data)
}

// Update
// @Summary		Update
// @Description	Update
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/file [Put]
func (h *FileController) Update(c echo.Context) error {
	var req dto.UpdateFileRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	_, err := h.fileService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)

	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

// Delete
// @Summary		Delete
// @Description	Delete
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/file [Delete]
func (h *FileController) Delete(c echo.Context) error {
	var req dto.DeleteFileRequest
	req.Id = utils.ParseStringIDFromUri(c)
	req.UserId = utils.GetUserStringIDFromContext(c)

	err := h.fileService.Delete(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

// GetOne
// @Summary		GetOne
// @Description	GetOne
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/file/download/:id [GET]
func (h *FileController) GetOne(c echo.Context) error {
	var req dto.GetOneFileRequest
	req.Id = utils.ParseStringIDFromUri(c)

	res, err := h.fileService.GetOne(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// Download
// @Summary		Download
// @Description	Download
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/file/download/:id [GET]
func (h *FileController) Download(c echo.Context) error {
	var req dto.DownloadFileRequest
	req.Id = utils.ParseStringIDFromUri(c)

	data, err := h.fileService.Download(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}
	res := data.Path + data.ID.String() + "." + data.ExtensionName

	return c.File(res)
}
