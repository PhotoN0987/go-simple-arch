package api

import (
	"go-simple-arch/pkg/model"

	"github.com/gin-gonic/gin"
)

// FileAPI file api
type FileAPI struct{}

// NewFileAPI file api create
func NewFileAPI(router *gin.RouterGroup) {
	fileAPI := FileAPI{}
	fileUploadRoutes := router.Group("/file")
	{
		fileUploadRoutes.GET("/:id", fileAPI.GetFile)
		fileUploadRoutes.POST("", fileAPI.UploadFile)
	}
}

// GetFile ファイルデータを取得します。
func (api *FileAPI) GetFile(c *gin.Context) {
	c.JSON(501, "開発中、できるまで待っててね。")
}

// UploadFile ファイルデータをUploadします。
func (api *FileAPI) UploadFile(c *gin.Context) {
	var file model.File
	c.BindJSON(&file)
	err := file.Create()
	if err == nil {
		c.JSON(201, "ファイルアップロード完了")
	} else {
		c.JSON(400, "ファイルアップロードに失敗しました")
	}
}
