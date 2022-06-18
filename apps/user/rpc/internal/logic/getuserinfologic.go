package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"lixiandea.github.com/go-zero-mall/apps/user/model"

	"lixiandea.github.com/go-zero-mall/apps/user/rpc/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *rpc.UserInfoRequest) (*rpc.UserInfoResponse, error) {
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.ID)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &rpc.UserInfoResponse{
		ID:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil

}
