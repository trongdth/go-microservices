package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trongdth/go_microservices/front-controller/serializers"
)

// DefaultWelcome : ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "API Gateway")
}

// Welcome : ...
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, serializers.Resp{Result: "REST API"})
}
