package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/logic"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/types"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
