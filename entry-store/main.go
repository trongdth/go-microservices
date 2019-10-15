package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/trongdth/go_microservices/entry-store/config"
	"github.com/trongdth/go_microservices/entry-store/daos"
	"github.com/trongdth/go_microservices/entry-store/servers"
	pb "github.com/trongdth/go_protobuf/entry-store"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conf := config.GetConfig()
	if err := daos.Init(conf); err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserSrvServer(grpcServer, servers.NewUserServer(
		daos.NewUser(),
	))
	grpcServer.Serve(lis)
}
