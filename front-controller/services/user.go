package services

import (
	"context"

	"github.com/pkg/errors"
	"github.com/trongdth/go_microservices/front-controller/config"
	errHandle "github.com/trongdth/go_microservices/front-controller/errors"
	"github.com/trongdth/go_microservices/front-controller/models"
	"github.com/trongdth/go_microservices/front-controller/serializers"
	"google.golang.org/grpc"

	pb "github.com/trongdth/go_protobuf"
	"golang.org/x/crypto/bcrypt"
)

// User : struct
type User struct {
	conf *config.Config
}

// NewUserService : config
func NewUserService(conf *config.Config) *User {
	return &User{
		conf: conf,
	}
}

// FindByID : ID
func (u *User) FindByID(ID uint) (*models.User, error) {
	conn, err := grpc.Dial(u.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, errors.Wrap(err, "u.Dial").Error())
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	r, err := c.ReadUser(context.Background(), &pb.UserReq{
		Req: &pb.BaseReq{
			Action:     pb.Action_QUERY,
			Message:    pb.Message_QUERY_USER_ID,
			ObjectType: pb.Object_USER,
		},
		User: &pb.UserInfo{
			Id: uint32(ID),
		},
	})

	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, errors.Wrap(err, "u.ReadUser").Error())
	}

	user := &models.User{
		ID:       uint(r.GetUser().Id),
		Email:    r.GetUser().Email,
		FullName: r.GetUser().FullName,
	}

	return user, nil
}

// Authenticate : user login request
func (u *User) Authenticate(req *serializers.UserLoginReq) (*models.User, error) {
	conn, err := grpc.Dial(u.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, errors.Wrap(err, "u.Dial").Error())
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	r, err := c.ReadUser(context.Background(), &pb.UserReq{
		Req: &pb.BaseReq{
			Action:     pb.Action_QUERY,
			Message:    pb.Message_QUERY_USER_EMAIL,
			ObjectType: pb.Object_USER,
		},
		User: &pb.UserInfo{
			Email: req.Email,
		},
	})

	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, errors.Wrap(err, "u.ReadUser").Error())
	}

	user := &models.User{
		ID:       uint(r.GetUser().Id),
		Email:    r.GetUser().Email,
		FullName: r.GetUser().FullName,
	}

	if err := bcrypt.CompareHashAndPassword([]byte(r.GetUser().Password), []byte(req.Password)); err != nil {
		return nil, errHandle.ErrInvalidPassword
	}

	return user, nil
}

// Register user: full name, email, password, confirm password
func (u *User) Register(req *serializers.UserRegisterReq) (*models.User, error) {

	if req.FullName == "" {
		return nil, errHandle.ErrInvalidName
	}

	if req.Email == "" {
		return nil, errHandle.ErrInvalidEmail
	}

	if req.Password == "" || req.ConfirmPassword == "" {
		return nil, errHandle.ErrInvalidPassword
	}

	if req.Password != req.ConfirmPassword {
		return nil, errHandle.ErrPasswordMismatch
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, "bcrypt.GenerateFromPassword")
	}

	conn, err := grpc.Dial(u.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, errors.Wrap(err, "u.Dial").Error())
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	r, err := c.CreateUser(context.Background(), &pb.UserReq{
		Req: &pb.BaseReq{
			Message:    pb.Message_STORE_CACHE_DB,
			ObjectType: pb.Object_USER,
			Action:     pb.Action_STORE,
		},
		User: &pb.UserInfo{
			FullName: req.FullName,
			Email:    req.Email,
			Password: string(hashed),
		},
	})

	if err != nil {
		return nil, errHandle.ErrorWithMessage(errHandle.ErrSystemError, errors.Wrap(err, "u.CreateUser").Error())
	}

	return &models.User{
		ID:       uint(r.GetUser().Id),
		Email:    r.GetUser().Email,
		FullName: r.GetUser().FullName,
	}, nil
}
