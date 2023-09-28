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

// SaveFile
//
// @Security		Authorization
// @Summary		SaveFile
// @Description	SaveFile
// @Tags			File
// @Accept			multipart/form-data
// @Param			file_request	formData	file		true	"file_request"
// @Success		200				{object}	model.File	"success"
// @Router			/v1/api/files [POST]
func (h *FileController) SaveFile(c echo.Context) error {
	file, err := c.FormFile("file_request")
	if err := c.Validate(file); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidationError(c, err)
	}

	req := request.UploadFileRequest{
		File:   file,
		UserId: userId,
	}
	data, err := h.fileService.SaveFile(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", data)
}

// Update
//
//	@Security		Authorization
//	@Summary		Update
//	@Description	Update
//	@Tags			File
//	@Accept			multipart/form-data
//	@Param			UpdateFileRequest	formData	request.UpdateFileRequest	true	"UpdateFileRequest"
//	@Param			file_request		formData	file						false	"file_request"
//	@Success		200					{object}	model.File					"success"
//	@Router			/v1/api/files [PUT]
func (h *FileController) Update(c echo.Context) error {
	var req request.UpdateFileRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}
	if err := c.Validate(req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidationError(c, err)
	}
	req.UserId = userId

	file, _ := c.FormFile("file_request")
	if file != nil {
		req.File = file
	}

	if err := c.Validate(req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	res, err := h.fileService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)

	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// Delete
//
//	@Security		Authorization
//	@Summary		Delete
//	@Description	Delete
//	@Tags			File
//	@Accept			json
//	@Param			id	path		string					true	"id"
//	@Success		200	{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/files [Delete]
func (h *FileController) Delete(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)

	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidationError(c, err)
	}
	req := request.DeleteFileRequest{
		ID:     id,
		UserId: userId,
	}

	if err := h.fileService.Delete(c.Request().Context(), req); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

// GetOne
//
//	@Security		Authorization
//	@Summary		GetOne
//	@Description	GetOne
//	@Tags			File
//	@Accept			json
//	@Param			id	path		string		true	"id"
//	@Success		200	{object}	model.File	"success"
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
//	@Security		Authorization
//	@Summary		Download
//	@Description	Download
//	@Tags			File
//	@Accept			json
//	@Param			id	path		string					true	"id"
//	@Success		200	{object}	entity.SimpleResponse	"success"
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
