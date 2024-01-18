package main

import (
	"context"

	pb "github.com/hemanth-ks97/grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello from the server",
	}, nil
}