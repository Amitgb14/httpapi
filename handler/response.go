package handler

type Response struct {
	Status int
	Body string
}

func (resp *Response) GetStatus() int {
	return resp.Status
}

func (resp *Response) GetBody() string {
	return resp.Body
}

