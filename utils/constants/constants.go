package constants

const (
	Local = "local"
	Dev   = "dev"
	Prod  = "prod"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

type TokenType string

const (
	AccessToken        TokenType = "access_token"
	RefreshToken       TokenType = "refresh_token"
	VerifyToken        TokenType = "verify_token"
	ResetPasswordToken TokenType = "reset_password_token"
)
