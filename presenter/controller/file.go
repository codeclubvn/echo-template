package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"trial_backend/pkg/utils"
	"trial_backend/presenter/request"
	"trial_backend/usecase"
)

type FileController struct {
	BaseController
	fileService usecase.FileService
}

func NewFileController(fileService usecase.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

// Upload
//
//	@Summary		Upload
//	@Description	Upload
//	@Tags			File
//	@Accept			multipart/form-data
//	@Param			Authorization	header		string		true	"authorization token"
//	@Param			file_request	formData	file		true	"file_request"
//	@Success		200				{object}	model.File	"success"
//	@Router			/v1/api/files [POST]
func (h *FileController) Upload(c echo.Context) error {
	file, err := c.FormFile("file_request")
	if err := c.Validate(file); err != nil {
		return h.ResponseValidationError(c, err)
	}

	data, err := h.fileService.Upload(c.Request().Context(), file)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", data)
}

// Update
//
//	@Summary		Update
//	@Description	Update
//	@Tags			File
//	@Accept			multipart/form-data
//	@Param			Authorization		header		string						true	"authorization token"
//	@Param			UpdateFileRequest	body		request.UpdateFileRequest	true	"UpdateFileRequest"
//	@Param			file_request		formData	file						true	"file_request"
//	@Success		200					{object}	model.File					"success"
//	@Router			/v1/api/files [Put]
func (h *FileController) Update(c echo.Context) error {
	var req request.UpdateFileRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	file, err := c.FormFile("file_request")
	if err != nil {
		return h.ResponseError(c, err)
	}

	if err := c.Validate(req); err != nil {
		return h.ResponseValidationError(c, err)
	}
	req.File = file

	res, err := h.fileService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)

	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// Delete
//
//	@Summary		Delete
//	@Description	Delete
//	@Tags			File
//	@Accept			json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Param			id				path		string					true	"id"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/files [Delete]
func (h *FileController) Delete(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	err := h.fileService.Delete(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

// GetOne
//
//	@Summary		GetOne
//	@Description	GetOne
//	@Tags			File
//	@Accept			json
//	@Param			Authorization	header		string		true	"authorization token"
//	@Param			id				path		string		true	"id"
//	@Success		200				{object}	model.File	"success"
//	@Router			/v1/api/files/{id} [GET]
func (h *FileController) GetOne(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	res, err := h.fileService.GetOne(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// Download
//
//	@Summary		Download
//	@Description	Download
//	@Tags			File
//	@Accept			json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Param			id				path		string					true	"id"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/files/download/{id} [GET]
func (h *FileController) Download(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	data, err := h.fileService.Download(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}
	res := data.Path + data.ID.String() + "." + data.ExtensionName

	return c.File(res)
}
