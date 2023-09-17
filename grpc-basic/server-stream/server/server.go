package main

import (
	"fmt"
	"github.com/hemantkgupta/my-go-work/grpc-basic/server-stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	proto.StreamServiceServer
}

func (s server) FetchResponse(in *proto.Request, srv proto.StreamService_FetchResponseServer) error {

	log.Printf("Sending response for request id : %d", in.Id)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		r := proto.Response{Result: fmt.Sprintf("Request Id:%d and response #%d", in.Id, i)}

		if err := srv.Send(&r); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("Sent response #%d", i)
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error in listening: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterStreamServiceServer(s, server{})

	log.Println("Server started.")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error in serving: %v", err)
	}

}
