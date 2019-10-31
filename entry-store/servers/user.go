package servers

import (
	"context"

	"github.com/trongdth/go_microservices/entry-store/daos"
	pb "github.com/trongdth/go_protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User : struct
type User struct {
	pb.UnimplementedUserSrvServer

	ud *daos.User
}

// NewUserServer :
func NewUserServer(ud *daos.User) *User {
	return &User{
		ud: ud,
	}
}

func (u *User) CreateUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (u *User) ReadUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}
