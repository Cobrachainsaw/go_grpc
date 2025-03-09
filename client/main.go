package main

import (
	"log"

	pb "github.com/Cobrachainsaw/go_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Name: []string{"Vin", "Alice", "Bob"},
	}

	//callSayHello(client)
	callSayHelloServerStream(client, names)
}
