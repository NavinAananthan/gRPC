package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"server_stream/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewCountDownClient(conn)

	timer := int32(10)

	stream, err := c.Start(context.Background(), &pb.CountdownRequest{Timer: timer})
	if err != nil {
		log.Fatalf("failed to start timer: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream read failed: %v", err)
		}

		fmt.Println(msg)
	}
}