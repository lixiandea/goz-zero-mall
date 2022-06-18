package main

import (
	"flag"
	"fmt"

	"lixiandea.github.com/go-zero-mall/apps/product/rpc/internal/config"
	"lixiandea.github.com/go-zero-mall/apps/product/rpc/internal/server"
	"lixiandea.github.com/go-zero-mall/apps/product/rpc/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewProductServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		product.RegisterProductServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
