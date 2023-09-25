package controller

import (
	"net/http"
	"trail_backend/dto"
	"trail_backend/utils"
	"trail_backend/utils/constants"

	"github.com/labstack/echo/v4"
	"trail_backend/service"
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

func (h *FileController) Create(c echo.Context) error {
	var req dto.UploadFileRequest
	err := utils.GetFile(c, &req, constants.FolderUpload)

	file, err := h.fileService.Create(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", file.ID)
}

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

func (h *FileController) GetList(c echo.Context) error {
	var req dto.GetListFileRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	files, err := h.fileService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", files)
}

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

func (h *FileController) Download(c echo.Context) error {
	var req dto.GetOneFileRequest
	req.Id = utils.ParseStringIDFromUri(c)

	res, err := h.fileService.GetOne(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}
