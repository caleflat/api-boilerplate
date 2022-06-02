package errors

import (
	"encoding/json"
	"net/http"
)

type ClientError interface {
	Error() string

	ResponseBody() ([]byte, error)

	ResponseHeaders() (int, map[string]string)
}

type HttpError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HttpError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}

	return e.Detail + ": " + e.Cause.Error()
}

func (e *HttpError) Write(rw http.ResponseWriter) {
	rw.WriteHeader(e.Status)
	body, err := e.ToJSON()
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write(body)
	rw.Header().Set("Content-Type", "application/json")
}

func (e *HttpError) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func NewHttpError(err error, status int, detail string) error {
	return &HttpError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}
