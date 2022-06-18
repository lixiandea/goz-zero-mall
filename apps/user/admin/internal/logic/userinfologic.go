package logic

import (
	"context"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/rpc"

	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &rpc.UserInfoRequest{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		ID:     res.ID,
		Gender: res.Gender,
		Name:   res.Name,
		Mobile: res.Mobile,
	}, nil
}
