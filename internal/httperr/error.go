package httperr

import (
	"fmt"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HTTPError struct {
	Errors map[string][]string `json:"errors"`
}

func NewHttpError(code int, reason, msg string) *HTTPError {
	fields := make(map[string][]string)
	fields[reason] = []string{msg}
	return &HTTPError{Errors: fields}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError message: %s", e.Errors)
}

// FromError try to convert an error to *HTTPError.
// 错误码格式要求： https://realworld-docs.netlify.app/specifications/backend/error-handling/
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}

	// 将框架内部错误，转换为合适的格式
	if se := new(errors.Error); errors.As(err, &se) {
		return NewHttpError(int(se.Code), se.Reason, se.Message)
	}

	// 业务错误
	return NewHttpError(442, "internal", err.Error())
}

func ErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(nethttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(nethttp.StatusInternalServerError)
	_, _ = w.Write(body)
}
