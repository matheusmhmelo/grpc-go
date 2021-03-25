package main

import (
	"context"
	"github.com/matheusmhmelo/grpc-go/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connet: %v", err)
	}

	defer connection.Close()

	Hello(connection, err)
}

func Hello(connection *grpc.ClientConn, err error) {
	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "Matheus",
	}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("error during the execution: %v", err)
	}

	log.Println(res.Msg)
}