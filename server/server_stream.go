package main

import (
	"log"
	"time"

	pb "github.com/Cobrachainsaw/go_grpc/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("got request with names: %v", req.Name)
	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Msg: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
