package main

import (
	"context"
	"fmt"
	"github.com/hemantkgupta/my-go-work/grpc-basic/chat/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	myChatServer := &server{}

	grpcServer := grpc.NewServer()
	proto.RegisterChatServiceServer(grpcServer, myChatServer)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error in serving: ", err)
	}

}

type server struct {
	proto.ChatServiceServer
}

func (s *server) SayHello(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	fmt.Println("Serving the hello")
	fmt.Println("Message from client: ", in.Body)
	return &proto.Message{Body: "Hello From the Server!"}, nil
}
