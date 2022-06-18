package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lixiandea.github.com/go-zero-mall/apps/product/admin/internal/logic"
	"lixiandea.github.com/go-zero-mall/apps/product/admin/internal/svc"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
