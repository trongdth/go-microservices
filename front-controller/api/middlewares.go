package api

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/trongdth/go_microservices/front-controller/serializers"

	"github.com/gin-gonic/gin"
	errHandle "github.com/trongdth/go_microservices/front-controller/errors"
	"github.com/trongdth/go_microservices/front-controller/models"
)

const ()

const (
	userIDKey    = "id"
	userEmailKey = "email"
	userRoleID   = "roleID"
)

// AuthMiddleware : ...
func AuthMiddleware(key string, authenticator func(c *gin.Context) (*models.User, error)) *jwt.GinJWTMiddleware {
	mw, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(key),
		Timeout:     1000 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: userIDKey,
		TokenLookup: "header:Authorization,query:token, cookie: jwt",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					userIDKey:    v.ID,
					userEmailKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := authenticator(c)

			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return user, nil
		},
		HTTPStatusMessageFunc: func(err error, c *gin.Context) string {
			return err.Error()
		},
		Unauthorized: func(c *gin.Context, _ int, message string) {
			c.JSON(http.StatusUnauthorized, serializers.Resp{
				Result: nil,
				Error:  errHandle.ErrorWithMessage(errHandle.ErrInvalidCredentials, message),
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, serializers.Resp{
				Result: serializers.UserLoginResp{
					Token:   token,
					Expired: expire.Format(time.RFC3339),
				},
				Error: nil,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, serializers.Resp{
				Result: serializers.UserLoginResp{
					Token:   token,
					Expired: expire.Format(time.RFC3339),
				},
				Error: nil,
			})
		},
	})
	return mw
}
