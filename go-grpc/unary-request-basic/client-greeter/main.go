package main

import (
	"context"
	"log"
	"time"

	pb "github.com/idiotcoffee/gRPC-learning/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	defaultName = "Ishaan"
	defaultAge  = 23
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect client: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetPersonClient(conn)

	cxt, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SayHi(cxt, &pb.HelloRequest{Name: defaultName, Age: int32(defaultAge)})

	if err != nil {
		log.Fatalf("error in getting response: %v", err)
	}
	log.Printf("Server returned: %s", response.GetWelcomeMsg())
}
