package constants

import "net/http"

var PublicRoutes = map[string]string{
	"/v1/api/auth/login":    http.MethodPost,
	"/v1/api/auth/register": http.MethodPost,
	"/v1/api/health/":       http.MethodGet,
	"/v1/api":               http.MethodPost,
}

const (
	NumberOfPath          = 3
	NumberOfJWTTokenArray = 2
)

const (
	GoogleUserInfoAPI = "https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token="
)
