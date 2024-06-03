package resp

import (
	"net/http"
	"time"
)

type Resp struct {
	Code      int    `json:"code"`
	Message   any    `json:"message"`
	IsError   bool   `json:"isError"`
	Err       string `json:"err"`
	Timestamp int64  `json:"ts"`
}

func New(code int, msg any, isErr bool, err string) Resp {
	return Resp{
		Code:      code,
		Message:   msg,
		IsError:   isErr,
		Err:       err,
		Timestamp: time.Now().Unix(),
	}
}

func Success(obj any) Resp {
	return New(http.StatusOK, obj, false, "")
}

func Fail(code int, err string) Resp {
	return New(code, nil, true, err)
}

func Error(err string) Resp {
	return New(http.StatusInternalServerError, nil, true, err)
}
