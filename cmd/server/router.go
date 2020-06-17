package server

import (
	"go-simple-arch/pkg/api"
	"go-simple-arch/pkg/repository"

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
	s.userRoutes(apiV1)
	return s.router
}

func (s *Server) userRoutes(rg *gin.RouterGroup) {
	repository := repository.NewUserRepository(s.db)
	api.NewUserAPI(rg, repository)
}
