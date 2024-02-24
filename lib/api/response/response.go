package response

const (
	StatusSuccess = "OK"
	StatusError   = "Error"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func OK() Response {
	return Response{
		Status: StatusSuccess,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
