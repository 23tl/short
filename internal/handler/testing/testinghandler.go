package testing

import (
	"go-zero-short/common/response"
	"net/http"

	"go-zero-short/internal/logic/testing"
	"go-zero-short/internal/svc"
)

func TestingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := testing.NewTestingLogic(r.Context(), svcCtx)
		resp, err := l.Testing()
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		response.Response(w, resp, err)
	}
}
