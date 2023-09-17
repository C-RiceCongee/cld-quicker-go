package v1

import (
	"cld-quicker/internal/api/v1/files"
	"cld-quicker/internal/api/v1/users"
	"github.com/gin-gonic/gin"
)

type IApiV1 interface {
	LaunchOnline(engine *gin.Engine)
}

type ApiV1 struct {
	files.IFileApi // 文件api
	users.IUserApi //用户api
}

// LaunchOnline ApiV1 启动上线的方法
func (ApiV1 *ApiV1) LaunchOnline(engine *gin.Engine) {
	file := engine.Group(ApiV1.IFileApi.SetPrefix())
	{
		file.GET("/list", ApiV1.IFileApi.GetFilesList())
	}
	user := engine.Group(ApiV1.IUserApi.SetPrefix())
	{
		user.GET("/list", ApiV1.IUserApi.GetUserList())
	}
}

// NewApiV1 api v1 接口初始化 内部初始化对应的模块构造代码
func NewApiV1() IApiV1 {
	return &ApiV1{
		IFileApi: files.NewFilesApi(),
		IUserApi: users.NewUsersApi(),
	}
}
