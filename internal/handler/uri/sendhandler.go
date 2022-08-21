package uri

import (
	"go-zero-short/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-short/internal/logic/uri"
	"go-zero-short/internal/svc"
	"go-zero-short/internal/types"
)

func SendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UrlRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := uri.NewSendLogic(r.Context(), svcCtx)
		resp, err := l.Send(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		response.Response(w, resp, err)
	}
}
