package main

import (
	"context"
	"fmt"
	"log"

	pb "pod-grpc/api/protos/v0"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequestv0{Name: "Alice", NestedData: &pb.NestedData1{Something1: "something"}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Greeting:", resp.GetMessage())
}
