package main

import (
	"context"
	"lets-learn-golang-grpc/chapter-00/helloworld/pb"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)

	// get a name otherwise defaultName
	name := defaultName
	if (len(os.Args)) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not say: %v", err)
	}
	log.Printf("Say: %s", r.GetMessage())
}
