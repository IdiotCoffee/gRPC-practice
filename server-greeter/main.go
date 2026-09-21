package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/idiotcoffee/gRPC-learning/greeter"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreetPersonServer
}

func (s *Server) SayHi(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: Name: %v, Age: %v", in.GetName(), in.GetAge())
	return &pb.HelloResponse{WelcomeMsg: fmt.Sprintf("hello, %v, are you %v years old?", in.GetName(), in.GetAge())}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprint(":50051"))
	if err != nil {
		log.Fatalf("Error listening to port: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterGreetPersonServer(serv, &Server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := serv.Serve(listen); err != nil {
		log.Fatalf("error in serving: %v", err)
	}
}
