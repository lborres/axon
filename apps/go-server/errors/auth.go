package errors

import "errors"

// Authentication Related Errors
var (
	// User Errors
	ErrUserExists         = errors.New("user already exists")       // 409
	ErrUserNotFound       = errors.New("user not found")            // 404
	ErrInvalidCredentials = errors.New("invalid email or password") // 401

	// Session Errors
	ErrInvalidToken         = errors.New("invalid session token")      // 401
	ErrSessionNotFound      = errors.New("session not found")          // 404
	ErrSessionExpired       = errors.New("session expired")            // 401
	ErrSessionCacheNotFound = errors.New("session not found in cache") // 404

	// Validation Errors
	ErrEmailRequired    = errors.New("email is required")     // 400
	ErrPasswordRequired = errors.New("password is required")  // 400
	ErrPasswordTooShort = errors.New("password is too short") // 400
	ErrPasswordTooLong  = errors.New("password is too long")  // 400
	ErrInvalidEmail     = errors.New("invalid email format")  // 400

	// Config Errors (internal, typically 500)
	ErrDBAdapterRequired   = errors.New("database adapter is required")          // 500
	ErrHTTPAdapterRequired = errors.New("http adapter is required")              // 500
	ErrSecretRequired      = errors.New("secret is required")                    // 500
	ErrSecretTooShort      = errors.New("secret must be at least 32 characters") // 500
)
