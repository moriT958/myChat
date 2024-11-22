package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrCode string

const (
	Unknown ErrCode = "U000"

	RenderHTMLFailed ErrCode = "U001"

	NoSessionFound      ErrCode = "A001"
	NoUserFound         ErrCode = "A002"
	CreateSessionFailed ErrCode = "A003"
	CreateUserFailed    ErrCode = "A004"
	DeleteSessionFailed ErrCode = "A005"

	ReadThreadFailed   ErrCode = "T001"
	CountRepliesFailed ErrCode = "T002"
	CreateThreadFailed ErrCode = "T003"
	CreatePostFailed   ErrCode = "T004"
	ReadPostFailed     ErrCode = "T005"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}

// Error that is shown to users.
type AppError struct {
	// enable to create error chain
	// by containing error in AppError.

	Err error `json:"-"` // don't show internal error.
	ErrCode
	Message string
}

// AppError implements error interface
func (appErr *AppError) Error() string {
	return appErr.Message
}

func (appErr *AppError) Unwrap() error {
	return appErr.Err
}

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *AppError

	// if unknown error occured.
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	// traceID := common.GetTraceID(req.Context())
	// log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	// Set responce code by error code.
	switch appErr.ErrCode {
	case NoSessionFound:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
