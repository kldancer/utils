package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "utils/grpc/hello-server/proto"
)

func main() {
	// 连接到Server端
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行RPC调用
	resp, err := client.SayHello(context.Background(), &pb.SayHelloRequest{RequestName: "ykl"})
	if err != nil {
		log.Errorf("failed to call: %v", err)
	}
	log.Infof("resp: %v", resp.GetResponseMsg())
}
