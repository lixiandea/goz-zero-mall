// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package server

import (
	"context"
	"lixiandea.github.com/go-zero-mall/apps/reply/rpc/internal/logic"
	"lixiandea.github.com/go-zero-mall/apps/reply/rpc/internal/svc"
	rpc2 "lixiandea.github.com/go-zero-mall/apps/reply/rpc/rpc"
)

type RpcServer struct {
	svcCtx *svc.ServiceContext
	rpc2.UnimplementedRpcServer
}

func NewRpcServer(svcCtx *svc.ServiceContext) *RpcServer {
	return &RpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcServer) Ping(ctx context.Context, in *rpc2.Request) (*rpc2.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
