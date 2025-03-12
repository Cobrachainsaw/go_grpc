package main

import (
	"io"
	"log"

	pb "github.com/Cobrachainsaw/go_grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStream(stream pb.GreetService_BiDirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Msg: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
