package servers

import (
	"context"

	"github.com/trongdth/go_microservices/entry-store/daos"
	pb "github.com/trongdth/go_protobuf/entry-store"
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

// GetUser :
func (u *User) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	id := req.GetId()

	user, err := u.ud.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	return &pb.UserRes{
		Id:       uint32(user.ID),
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.UserName,
	}, nil
}
