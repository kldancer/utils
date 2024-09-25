package main

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "utils/grpc/hello-server/proto"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (server) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建GRPC服务
	serv := grpc.NewServer()
	// 注册自己编写端服务
	pb.RegisterSayHelloServer(serv, &server{})

	// 启动服务
	if err := serv.Serve(listen); err != nil {
		log.Errorf("GRPC Server stopped: %v", err)
		panic(err)
	}
}
