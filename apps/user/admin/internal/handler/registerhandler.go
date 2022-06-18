package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/logic"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
