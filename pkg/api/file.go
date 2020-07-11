package api

import (
	"go-simple-arch/pkg/model"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// FileAPI file api
type FileAPI struct{}

// NewFileAPI file api create
func NewFileAPI(router *gin.RouterGroup) {
	fileAPI := FileAPI{}
	fileUploadRoutes := router.Group("/file")
	{
		fileUploadRoutes.GET("/:filename", fileAPI.GetFile)
		fileUploadRoutes.POST("", fileAPI.UploadFile)
	}
}

// GetFile ファイルデータをDownloadします。
func (api *FileAPI) GetFile(c *gin.Context) {
	directory := "web/upload_files/"

	fileName := c.Param("filename")
	targetPath := filepath.Join(directory, fileName)

	if !strings.HasPrefix(filepath.Clean(targetPath), directory) {
		c.String(403, "ファイルのダウンロードに失敗しました")
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}

// UploadFile ファイルデータをUploadします。
func (api *FileAPI) UploadFile(c *gin.Context) {
	var file model.File
	c.BindJSON(&file)
	err := file.Create()
	if err == nil {
		c.String(201, "ファイルアップロード完了")
	} else {
		c.String(400, "ファイルのアップロードに失敗しました")
	}
}
