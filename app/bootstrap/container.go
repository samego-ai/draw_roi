package bootstrap

import (
	"github.com/kataras/golog"
	"github.com/samego-ai/draw_roi/app/handler"
	"github.com/samego-ai/draw_roi/config"
	"github.com/samego-ai/draw_roi/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Container 业务实现方法的容器
type Container struct {
	Listen net.Listener
	Server *grpc.Server
}

// New 实例化微服务容器
func New() *Container {
	golog.Info(config.App.Protocol, config.App.Address)
	listen, err := net.Listen(config.App.Protocol, config.App.Address)
	if err != nil {
		log.Fatalf("server listen failure: %v", err)
	}
	return &Container{
		Listen: listen,
		Server: grpc.NewServer(),
	}
}

// Register 注册 grpc 服务
func (container *Container) Register() *Container {
	pb.RegisterDrawServer(container.Server, &handler.DrawServer{})

	return container
}

// Run 运行容器服务
func (container *Container) Run() {
	err := container.Server.Serve(container.Listen)
	if nil != err {
		log.Fatalf("server run failure: %v", err)
	}
}
