package main

import (
	"fmt"
	"log"
	"net"

	"github.com/trongdth/go_microservices/entry-cache/config"
	"github.com/trongdth/go_microservices/entry-cache/daos"
	"github.com/trongdth/go_microservices/entry-cache/servers"
	"github.com/trongdth/go_microservices/entry-cache/services"
	pb "github.com/trongdth/go_protobuf"
	"google.golang.org/grpc"
)

func main() {
	conf := config.GetConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := daos.Init(conf); err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserSrvServer(grpcServer, servers.NewUserServer(services.NewUserSvc(conf), daos.NewUser()))
	grpcServer.Serve(lis)
}
