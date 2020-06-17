package api

import (
	"database/sql"
	"go-simple-arch/pkg/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserAPI user api
type UserAPI struct {
	repository repository.UserRepository
}

// NewUserAPI user api create
func NewUserAPI(router *gin.RouterGroup, repo repository.UserRepository) {
	userAPI := &UserAPI{
		repository: repo,
	}
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userAPI.GetAll)
		userRoutes.GET("/:id", userAPI.GetByID)
	}
}

// GetAll 複数のUserを取得します
func (api *UserAPI) GetAll(c *gin.Context) {
	result, err := api.repository.GetAll()

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetByID UserをIDで検索して取得します
func (api *UserAPI) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := api.repository.GetByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
