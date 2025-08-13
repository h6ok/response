package response

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Writer    http.ResponseWriter `json:"-"`
	Status    int                 `json:"status"`
	Data      any                 `json:"data,omitempty"`
	Error     ErrorResponse       `json:"error"`
	Timestamp time.Time           `json:"timestamp"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func Success(r http.ResponseWriter) *Response {
	return &Response{
		Writer: r,
		Status: http.StatusOK,
	}
}

func Accepted(r http.ResponseWriter) *Response {
	return &Response{
		Writer: r,
		Status: http.StatusAccepted,
	}
}

func BadRequest(r http.ResponseWriter) *Response {
	return &Response{
		Writer: r,
		Status: http.StatusBadRequest,
	}
}

func ServerError(r http.ResponseWriter) *Response {
	return &Response{
		Writer: r,
		Status: http.StatusInternalServerError,
	}
}

func Status(r http.ResponseWriter, status int) *Response {
	return &Response{
		Writer: r,
		Status: status,
	}
}

func (res *Response) Json() *Response {
	res.Writer.Header().Set("Content-Type", "application/json")
	return res
}

func (res *Response) SetBody(data any) *Response {
	res.Data = data
	res.Timestamp = time.Now()
	return res
}

func (res *Response) BasicSecurity() *Response {
	res.Writer.Header().Set("X-Content-Type-Options", "nosniff")
	res.Writer.Header().Set("X-Frame-Options", "DENY")
	res.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
	res.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	return res
}

func (res *Response) CORS() *Response {
	res.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	res.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	res.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	res.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	return res
}

func (res *Response) SetError(err error) *Response {
	errRes := ErrorResponse{
		Message: err.Error(),
	}

	res.Error = errRes
	return res
}

func (res *Response) SetHeader(key, value string) *Response {
	res.Writer.Header().Set(key, value)
	return res
}

func (res *Response) SetHeaders(headers map[string]string) *Response {
	for k, v := range headers {
		res.Writer.Header().Set(k, v)
	}
	return res
}

func (res *Response) Return() {
	data, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	res.Writer.Write(data)
}
