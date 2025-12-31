package errors

import "errors"

// General HTTP Errors
var (
	ErrValidation         = errors.New("input validation failed")          // 400
	ErrUnauthorized       = errors.New("authentication required")          // 401
	ErrForbidden          = errors.New("insufficient permissions")         // 403
	ErrResourceNotFound   = errors.New("requested resource doesn't exist") // 404
	ErrConflict           = errors.New("resource conflict")                // 409
	ErrRateLimitExceeded  = errors.New("too many requests")                // 429
	ErrInternalServer     = errors.New("internal server error")            // 500
	ErrNotImplemented     = errors.New("not implemented")                  // 501
	ErrServiceUnavailable = errors.New("temporary service outage")         // 503
)
