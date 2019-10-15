package main

import (
	"context"
	"log"

	pb "github.com/trongdth/go_protobuf/entry-store"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserSrvClient(conn)

	r, err := c.GetUser(context.Background(), &pb.UserReq{Id: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Greeting: ", r.GetFullName(), r.GetEmail())
}
