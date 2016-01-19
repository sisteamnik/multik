package multik

import (
	"net/http"
)

type Response struct {
	Status      int
	ContentType string

	Out http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{Out: w}
}
