package services

import (
	"context"

	"github.com/pkg/errors"
	"github.com/trongdth/go_microservices/entry-cache/config"
	pb "github.com/trongdth/go_protobuf"
	"google.golang.org/grpc"
)

// UserSvc : struct
type UserSvc struct {
	conf *config.Config
}

// NewUserSvc :
func NewUserSvc(conf *config.Config) *UserSvc {
	return &UserSvc{
		conf: conf,
	}
}

// ReadUser : context, user request
func (urv *UserSvc) ReadUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	conn, err := grpc.Dial(urv.conf.EntryStoreEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "u.Dial")
	}
	defer conn.Close()
	pbEntryStore := pb.NewUserSrvClient(conn)
	return pbEntryStore.ReadUser(context.Background(), req)
}

// CreateUser :
func (urv *UserSvc) CreateUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	conn, err := grpc.Dial(urv.conf.EntryStoreEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "u.Dial")
	}
	defer conn.Close()
	pbEntryStore := pb.NewUserSrvClient(conn)
	return pbEntryStore.CreateUser(context.Background(), req)
}
