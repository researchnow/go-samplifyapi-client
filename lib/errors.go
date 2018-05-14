package samplify

import (
	"fmt"
	"strings"
	"time"
)

// Error ...
type Error struct {
	Path         string `json:"path"`
	Message      string `json:"message"`
	InvalidValue string `json:"invalidValue"`
	Constraint   string `json:"constraint"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Timestamp  *time.Time `json:"timestamp"`
	RequestID  string     `json:"requestId"`
	Path       string     `json:"path"`
	HTTPCode   int        `json:"httpCode"`
	HTTPPhrase int        `json:"httpPhrase"`
	Errors     []*Error   `json:"errors"`
}

// Error ...
func (e *ErrorResponse) Error() string {
	str := ""
	for _, err := range e.Errors {
		str = fmt.Sprintf("%s\n%s", str, err.Message)
	}
	return strings.TrimSpace(str)
}
