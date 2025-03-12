package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Cobrachainsaw/go_grpc/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.BiDirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()
}
