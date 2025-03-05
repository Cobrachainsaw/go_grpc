package main

import (
	"context"

	pb "github.com/Cobrachainsaw/go_grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: "Hello",
	}, nil
}
