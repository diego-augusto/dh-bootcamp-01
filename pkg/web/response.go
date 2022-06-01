package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}) Response {
	return Response{code, data, ""}
}

func DecodeError(code int, err string) Response {
	return Response{code, nil, err} // Erro
}
