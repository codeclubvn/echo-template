package api_errors

import "net/http"

var (
	ErrInternalServerError  = "10000"
	ErrUnauthorizedAccess   = "10001"
	ErrTokenBadSignedMethod = "10002"
	ErrTokenExpired         = "10003"
	ErrTokenInvalid         = "10004"
	ErrTokenMalformed       = "10005"
	ErrUserNotFound         = "10006"
	ErrRequestTimeout       = "10008"
	ErrTokenMissing         = "10009"
	ErrValidation           = "10010"
	ErrInvalidUserID        = "10011"
	ErrMissingXStoreID      = "10012"
	ErrPermissionDenied     = "10013"
	ErrInvalidPassword      = "10014"
	CannotParseToken        = "10015"
	ErrPostNotFound         = "10016"
	ErrFileNotFound         = "10017"
	ErrEmailExist           = "10018"
	ListFilesInvalid        = "10019"
	RequestInvalid          = "10020"
	FileIsNotImage          = "10021"
	ErrFileTooLarge         = "10022"
)

type MessageAndStatus struct {
	Message string
	Status  int
}

var MapErrorCodeMessage = map[string]MessageAndStatus{
	ErrInternalServerError:  {"Internal Server Error", http.StatusInternalServerError},
	ErrUnauthorizedAccess:   {"Unauthorized Access", http.StatusUnauthorized},
	ErrTokenBadSignedMethod: {"Token Bad Signed Method", http.StatusUnauthorized},
	ErrTokenExpired:         {"Token Expired", http.StatusUnauthorized},
	ErrTokenInvalid:         {"Token Invalid", http.StatusUnauthorized},
	ErrTokenMalformed:       {"Token Malformed", http.StatusUnauthorized},
	ErrUserNotFound:         {"User Not Found", http.StatusNotFound},
	ErrPostNotFound:         {"Post Not Found", http.StatusNotFound},
	ErrFileNotFound:         {"File Not Found", http.StatusNotFound},
	ErrRequestTimeout:       {"Request Timeout", http.StatusRequestTimeout},
	ErrTokenMissing:         {"Token Missing", http.StatusUnauthorized},
	ErrValidation:           {"Validation Error", http.StatusBadRequest},
	ErrInvalidUserID:        {"Invalid User ID", http.StatusBadRequest},
	ErrMissingXStoreID:      {"Missing x-store-id", http.StatusBadRequest},
	ErrPermissionDenied:     {"Permission Denied", http.StatusForbidden},
	ErrInvalidPassword:      {"Invalid Password", http.StatusBadRequest},
	CannotParseToken:        {"Cannot Parse Token", http.StatusUnauthorized},
	ErrEmailExist:           {"Email is Exist", http.StatusBadRequest},
	ListFilesInvalid:        {"files field must contain a UUID", http.StatusUnprocessableEntity},
	RequestInvalid:          {"Request invalid", http.StatusBadRequest},
	FileIsNotImage:          {"Only accept type image: png, jpg, jpeg, gif, svg", http.StatusBadRequest},
	ErrFileTooLarge:         {"File size must not exceed 25MB", http.StatusBadRequest},
}
