package files

import (
	"cld-quicker/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IFileApi interface {
	api.IBase
	GetFilesList() func(ctx *gin.Context)
}

type FileApi struct {
}

func (f *FileApi) SetPrefix() string {
	return "/api/v1/files"
}
func (f *FileApi) GetFilesList() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "获取文件列表",
		})
	}
}
func NewFilesApi() IFileApi {
	return &FileApi{}
}
