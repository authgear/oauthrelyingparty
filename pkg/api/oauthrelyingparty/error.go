package oauthrelyingparty

import (
	"strings"
)

type ErrorCategory string

const (
	ErrorCategoryCancelled ErrorCategory = "cancelled"
	ErrorCategoryRejected  ErrorCategory = "rejected"
	ErrorCategoryFailed    ErrorCategory = "failed"
	ErrorCategoryTimeout   ErrorCategory = "timeout"
)

type ErrorResponse struct {
	Error_           string        `json:"error"`
	ErrorDescription string        `json:"error_description,omitempty"`
	ErrorURI         string        `json:"error_uri,omitempty"`
	Category         ErrorCategory `json:"-"`
}

func (e *ErrorResponse) Error() string {
	var buf strings.Builder
	buf.WriteString(e.Error_)
	if e.ErrorDescription != "" {
		buf.WriteString(": ")
		buf.WriteString(e.ErrorDescription)
	}
	return buf.String()
}
