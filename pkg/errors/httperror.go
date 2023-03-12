package errors

type HttpError struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}
