package zex

import "net/http"

// Error is a custom error type
type Error struct {
	status   int
	message  string
	internal error
}

// NewError creates a new error
func NewError(status int, message string) *Error {
	return &Error{
		status:   status,
		message:  message,
		internal: nil,
	}
}

// Error returns the error message
func (e *Error) Error() string {
	return e.message
}

// Status returns the error status
func (e *Error) Status() int {
	return e.status
}

// SetInternal sets the internal error
func (e *Error) SetInternal(err error) {
	e.internal = err
}

// Internal returns the internal error
func (e *Error) Internal() error {
	return e.internal
}

var (
	// Client Errors 4xx

	ErrBadRequest                 = NewError(http.StatusBadRequest, "Bad Request")
	ErrUnauthorized               = NewError(http.StatusUnauthorized, "Unauthorized")
	ErrPaymentRequired            = NewError(http.StatusPaymentRequired, "Payment Required")
	ErrForbidden                  = NewError(http.StatusForbidden, "Forbidden")
	ErrNotFound                   = NewError(http.StatusNotFound, "Not Found")
	ErrMethodNotAllowed           = NewError(http.StatusMethodNotAllowed, "Method Not Allowed")
	ErrNotAcceptable              = NewError(http.StatusNotAcceptable, "Not Acceptable")
	ErrProxyAuthRequired          = NewError(http.StatusProxyAuthRequired, "Proxy Authentication Required")
	ErrRequestTimeout             = NewError(http.StatusRequestTimeout, "Request Timeout")
	ErrConflict                   = NewError(http.StatusConflict, "Conflict")
	ErrGone                       = NewError(http.StatusGone, "Gone")
	ErrLengthRequired             = NewError(http.StatusLengthRequired, "Length Required")
	ErrPreconditionFailed         = NewError(http.StatusPreconditionFailed, "Precondition Failed")
	ErrPayloadTooLarge            = NewError(http.StatusRequestEntityTooLarge, "Payload Too Large")
	ErrURITooLong                 = NewError(http.StatusRequestURITooLong, "URI Too Long")
	ErrUnsupportedMedia           = NewError(http.StatusUnsupportedMediaType, "Unsupported Media Type")
	ErrRangeNotSatisfiable        = NewError(http.StatusRequestedRangeNotSatisfiable, "Range Not Satisfiable")
	ErrExpectationFailed          = NewError(http.StatusExpectationFailed, "Expectation Failed")
	ErrTeapot                     = NewError(http.StatusTeapot, "I'm a teapot")
	ErrMisdirectedRequest         = NewError(http.StatusMisdirectedRequest, "Misdirected Request")
	ErrUnprocessableEntity        = NewError(http.StatusUnprocessableEntity, "Unprocessable Entity")
	ErrLocked                     = NewError(http.StatusLocked, "Locked")
	ErrFailedDependency           = NewError(http.StatusFailedDependency, "Failed Dependency")
	ErrTooEarly                   = NewError(http.StatusTooEarly, "Too Early")
	ErrUpgradeRequired            = NewError(http.StatusUpgradeRequired, "Upgrade Required")
	ErrPreconditionRequired       = NewError(http.StatusPreconditionRequired, "Precondition Required")
	ErrTooManyRequests            = NewError(http.StatusTooManyRequests, "Too Many Requests")
	ErrHeaderFieldsTooLarge       = NewError(http.StatusRequestHeaderFieldsTooLarge, "Request Header Fields Too Large")
	ErrUnavailableForLegalReasons = NewError(http.StatusUnavailableForLegalReasons, "Unavailable For Legal Reasons")

	// Server Errors 5xx

	ErrInternalServerError           = NewError(http.StatusInternalServerError, "Internal Server Error")
	ErrNotImplemented                = NewError(http.StatusNotImplemented, "Not Implemented")
	ErrBadGateway                    = NewError(http.StatusBadGateway, "Bad Gateway")
	ErrServiceUnavailable            = NewError(http.StatusServiceUnavailable, "Service Unavailable")
	ErrGatewayTimeout                = NewError(http.StatusGatewayTimeout, "Gateway Timeout")
	ErrHTTPVersionNotSupported       = NewError(http.StatusHTTPVersionNotSupported, "HTTP Version Not Supported")
	ErrVariantAlsoNegotiates         = NewError(http.StatusVariantAlsoNegotiates, "Variant Also Negotiates")
	ErrInsufficientStorage           = NewError(http.StatusInsufficientStorage, "Insufficient Storage")
	ErrLoopDetected                  = NewError(http.StatusLoopDetected, "Loop Detected")
	ErrNotExtended                   = NewError(http.StatusNotExtended, "Not Extended")
	ErrNetworkAuthenticationRequired = NewError(http.StatusNetworkAuthenticationRequired, "Network Authentication Required")
)
