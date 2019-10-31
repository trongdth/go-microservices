package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/trongdth/go_microservices/entry-store/config"
	"github.com/trongdth/go_microservices/entry-store/daos"
	"github.com/trongdth/go_microservices/entry-store/servers"
	pb "github.com/trongdth/go_protobuf"
	"google.golang.org/grpc"
)

func main() {
	conf := config.GetConfig()
	if err := daos.Init(conf); err != nil {
		panic(err)
	}

	if err := daos.AutoMigrate(); err != nil {
		log.Fatal("failed to auto migrate", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserSrvServer(grpcServer, servers.NewUserServer(
		daos.NewUser(),
	))
	grpcServer.Serve(lis)
}
