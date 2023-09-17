package main

import (
	"context"
	"github.com/hemantkgupta/my-go-work/grpc-basic/server-stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {

	con, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error in dialing: %v", err)
	}

	c := proto.NewStreamServiceClient(con)
	in := &proto.Request{Id: 1}
	stream, err := c.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("Error in fetching response: %v", err)
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error in response: %v", err)
		}
		log.Printf("Response received for: %s", r.Result)
	}
}
