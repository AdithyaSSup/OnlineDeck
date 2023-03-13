package errors

import "fmt"

type HttpError struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (e HttpError) Error() string {
	msg := fmt.Sprintf("ERROR: %s: %s: %s", e.Type, e.Title, e.Detail)
	return msg
}
