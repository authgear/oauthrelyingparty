package oauthrelyingparty

import (
	"strings"
)

type ErrorResponse struct {
	Error_           string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
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
