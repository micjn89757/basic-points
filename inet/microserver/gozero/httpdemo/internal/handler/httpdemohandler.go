package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"inet/microserver/gozero/httpdemo/internal/logic"
	"inet/microserver/gozero/httpdemo/internal/svc"
	"inet/microserver/gozero/httpdemo/internal/types"
)

func HttpdemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHttpdemoLogic(r.Context(), svcCtx)
		resp, err := l.Httpdemo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
