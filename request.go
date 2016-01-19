package multik

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
)

type Request struct {
	*http.Request
	ContentType string
	Format      string // "html", "xml", "json", or "txt"
	//AcceptLanguages AcceptLanguages
	Locale    string
	Websocket *websocket.Conn
}

func NewRequest(r *http.Request) *Request {
	return &Request{Request: r}
}
