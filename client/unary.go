package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hemanth-ks97/grpc-demo/proto"
)

func CallSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil{
		log.Fatalf("Could not greet, %v", err)
	}
	log.Printf("%s", res)
}