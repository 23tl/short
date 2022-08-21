package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-short/common/errorx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError:
			body.Code = e.Date().Code
			body.Message = e.Date().Message
		default:
			body.Code = -1
			body.Message = e.Error()
		}
	} else {
		body.Message = "ok"
		body.Data = resp
	}

	httpx.OkJson(w, body)
}
