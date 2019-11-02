package servers

import (
	"context"

	"github.com/trongdth/go_microservices/entry-cache/services"
	pb "github.com/trongdth/go_protobuf"
)

// User : struct
type User struct {
	userSvc *services.UserSvc
}

// NewUserServer :
func NewUserServer(userSvc *services.UserSvc) *User {
	return &User{
		userSvc: userSvc,
	}
}

// CreateUser : context, user request
func (u *User) CreateUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	return u.userSvc.CreateUser(ctx, req)
}

// ReadUser : context, user request
func (u *User) ReadUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {

	return u.userSvc.ReadUser(ctx, req)
}
