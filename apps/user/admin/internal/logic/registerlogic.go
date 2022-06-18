package logic

import (
	"context"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/rpc"

	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &rpc.RegisterRequest{
		Gender:   req.Gender,
		Name:     req.Name,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{
		Gender: res.Gender,
		Name:   res.Name,
		Mobile: res.Mobile,
		ID:     res.ID,
	}, nil
}
