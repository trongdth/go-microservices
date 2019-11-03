package servers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/trongdth/go_microservices/entry-cache/daos"
	"github.com/trongdth/go_microservices/entry-cache/services"
	pb "github.com/trongdth/go_protobuf"
)

// User : struct
type User struct {
	userSvc *services.UserSvc
	ud      *daos.User
}

// NewUserServer :
func NewUserServer(userSvc *services.UserSvc, ud *daos.User) *User {
	return &User{
		userSvc: userSvc,
		ud:      ud,
	}
}

// CreateUser : context, user request
func (u *User) CreateUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	return u.userSvc.CreateUser(ctx, req)
}

// ReadUser : context, user request
func (u *User) ReadUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	if req.Req.Message == pb.Message_QUERY_USER_ID {
		ID := req.User.GetId()
		cache, _ := u.ud.GetUser(ID)
		if cache != nil {
			var userInfo *pb.UserInfo
			err := json.Unmarshal([]byte(cache.(string)), &userInfo)
			if err == nil {
				fmt.Println("return data from cache")
				return &pb.UserRes{
					Result: true,
					User:   userInfo,
				}, nil
			}
		}
	}

	res, err := u.userSvc.ReadUser(ctx, req)
	if res != nil {
		err := u.ud.SetUser(res.GetUser().Id, res.GetUser())
		if err != nil {
			log.Println(err)
		}
	}

	return res, err
}
