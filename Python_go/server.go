package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"python_go/pb"
	"google.golang.org/grpc"
	
)

type server struct {
	pb.UnimplementedCommunicationServiceServer
}

func (s *server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	fmt.Println("Received message:", req.GetMessage())
	return &pb.MessageResponse{Response: "Hello from Go!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCommunicationServiceServer(s, &server{})

	fmt.Println("Server is listening on port 50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
