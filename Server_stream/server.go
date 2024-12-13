package main

import (
	"log"
	"net"
	"time"
	"fmt"

	"server_stream/pb"
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedCountDownServer
}


func main() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCountDownServer(s, &server{})

	fmt.Println("Server is listening on port 50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func (*server) Start(req *pb.CountdownRequest, stream pb.CountDown_StartServer) error {
	t := req.GetTimer()
	
	for t > 0 {
		res := pb.CountDownResponse{Count: t}
		stream.Send(&res)
		t--
		time.Sleep(time.Second)
	}
	return nil
}