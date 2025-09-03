package constants

const (
	ErrInvalidBody        = "Invalid request body"
	ErrInvalidCredentials = "Invalid email or password"
	ErrInternalServer     = "Internal server error"
	ErrTokenGeneration    = "Token generation failed"
	ErrUnauthorized       = "Unauthorized access"
	ErrForbidden          = "Forbidden"
	ErrNotFound           = "Resource not found"
	ErrBadRequest         = "Bad request"
	ErrValidationFailed   = "Validation failed"
	ErrUserNotFound       = "User not found"
	ErrUserAlreadyExists  = "User already exists"
	ErrDatabaseConnection = "Database connection failed"
	ErrDatabaseQuery      = "Database query failed"
)

// Reason codes for detailed error responses
const (
	ReasonInvalidCredentials    = "invalid_credentials"
	ReasonTokenGenerationFailed = "token_generation_failed"
	ReasonUnknownError          = "unknown_error"
	ReasonNilToken              = "nil_token"
)
