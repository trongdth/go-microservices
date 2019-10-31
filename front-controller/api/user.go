package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	errHandle "github.com/trongdth/go_microservices/front-controller/errors"
	"github.com/trongdth/go_microservices/front-controller/models"

	"github.com/trongdth/go_microservices/front-controller/serializers"
)

// Authenticate an user
func (s *Server) Authenticate(c *gin.Context) (*models.User, error) {
	var req serializers.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	user, err := s.userSvc.Authenticate(&req)
	if err != nil {
		return nil, errors.Wrap(err, "u.svc.Authenticate")
	}

	return user, nil
}

// Register an user
func (s *Server) Register(c *gin.Context) {
	var req serializers.UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: errHandle.ErrInvalidArgument})
		return
	}

	user, err := s.userSvc.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: err})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: user, Error: nil})
}

// UserProfile of user
func (s *Server) UserProfile(c *gin.Context) {
	user, err := s.userFromContext(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: errHandle.ErrInvalidCredentials})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: user, Error: nil})
}
