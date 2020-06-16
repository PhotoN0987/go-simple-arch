package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))

	return router
}

// SetUpRouter Setup all api routing
func (s *Server) SetUpRouter() *gin.Engine {
	// Group v1
	apiV1 := s.router.Group("api/v1")
	s.testRoutes(apiV1)
	return s.router
}

func (s *Server) testRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/test")
	{
		userRoutes.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, "test")
		})
	}
}
