package multik

type Result interface {
	Apply(req *Request, resp *Response)
}
