package base

type Response struct {
	Data    any    `json:"data"`
	Error   error  `json:"error"`
	Message string `json:"message"`
}

func NewResponse(data any, message string) *Response {
	return &Response{
		Data:    data,
		Message: message,
	}
}

func (r *Response) WithError(err error) *Response {
	r.Error = err
	return r
}
