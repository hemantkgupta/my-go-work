package main

import (
	"context"
	"fmt"
	"github.com/hemantkgupta/my-go-work/grpc-basic/chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error in dialing: ", err)
	}
	defer con.Close()

	cl := proto.NewChatServiceClient(con)

	res, err := cl.SayHello(context.Background(), &proto.Message{Body: "Hello World!"})

	if err != nil {
		fmt.Println("Error in getting response: ", err)
	}
	fmt.Println("The response from server: ", res.Body)
}
