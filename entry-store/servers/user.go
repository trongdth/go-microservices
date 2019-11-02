package servers

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	errHandle "github.com/tokoinofficial/entry_store_service/errors"
	"github.com/trongdth/go_microservices/entry-store/daos"
	"github.com/trongdth/go_microservices/entry-store/models"
	pb "github.com/trongdth/go_protobuf"
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

func (urv *User) isValidUserQueryMessage(req *pb.BaseReq) bool {
	if req.Action == pb.Action_QUERY &&
		req.ObjectType == pb.Object_USER &&
		(req.Message == pb.Message_QUERY_USER_ID || req.Message == pb.Message_QUERY_USER_EMAIL) {
		return true
	}
	return false
}

// CreateUser : ...
func (urv *User) CreateUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	email := req.GetUser().GetEmail()
	user, err := urv.ud.FindByEmail(email)
	if user != nil {
		return nil, errors.New("Email has existed already")
	}

	newUser := &models.User{
		FullName: req.GetUser().GetFullName(),
		Email:    req.GetUser().GetEmail(),
		Password: req.GetUser().GetPassword(),
	}

	err = daos.WithTransaction(func(tx *gorm.DB) error {
		err := urv.ud.Create(tx, newUser)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserRes{
		Result: true,
		User: &pb.UserInfo{
			Id:       uint32(newUser.ID),
			FullName: newUser.FullName,
			Email:    newUser.Email,
			Password: newUser.Password,
		},
	}, nil
}

// ReadUser : ...
func (urv *User) ReadUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	if !urv.isValidUserQueryMessage(req.Req) {
		return nil, errHandle.ErrInvalidMessage
	}

	var user *models.User
	var err error

	switch req.Req.Message {
	case pb.Message_QUERY_USER_ID:
		userID := req.User.GetId()
		user, err = urv.ud.FindByID(uint(userID))
		if err != nil {
			return nil, err
		}
	case pb.Message_QUERY_USER_EMAIL:
		email := req.User.GetEmail()
		user, err = urv.ud.FindByEmail(email)
		if err != nil {
			return nil, err
		}

	}

	return &pb.UserRes{
		Result: true,
		User: &pb.UserInfo{
			Id:       uint32(user.ID),
			FullName: user.FullName,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}
