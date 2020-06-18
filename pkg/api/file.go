package api

import "github.com/gin-gonic/gin"

// FileAPI file api
type FileAPI struct{}

// NewFileAPI file api create
func NewFileAPI(router *gin.RouterGroup) {
	fileUploadRoutes := router.Group("/file")
	{
		fileUploadRoutes.GET("/:fileID", nil)
		fileUploadRoutes.POST("", nil)
	}
}

// GetFile ファイルデータを取得します。
func (api *FileAPI) GetFile(c *gin.Context) {}

// UploadFile ファイルデータをUploadします。
func (api *FileAPI) UploadFile(c *gin.Context) {}
