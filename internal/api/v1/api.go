package v1

import (
	"cld-quicker-go/internal/api/v1/files"
	"cld-quicker-go/internal/api/v1/foods"
	"cld-quicker-go/internal/api/v1/users"

	"github.com/gin-gonic/gin"
)

type IApiV1 interface {
	LaunchOnline(engine *gin.Engine)
}

type ApiV1 struct {
	files.IFileApi //File API
	users.IUserApi //User API
	foods.IFoodsApi
}

// LaunchOnline ApiV1 Initiation of go-live methods
func (ApiV1 *ApiV1) LaunchOnline(engine *gin.Engine) {
	file := engine.Group(ApiV1.IFileApi.SetPrefix())
	{
		file.GET("/list", ApiV1.IFileApi.GetFilesList())
		file.POST("/singleUpload", ApiV1.IFileApi.SingleUploadFile())
	}
	user := engine.Group(ApiV1.IUserApi.SetPrefix())
	{
		user.GET("/list", ApiV1.IUserApi.GetUserList())
	}
	foods := engine.Group(ApiV1.IFoodsApi.SetPrefix())
	{
		foods.GET("/common", ApiV1.IFoodsApi.GetCommonFoods())
	}
}

// NewApiV1 api v1 Interface initialization Internal initialization corresponds to the module construction code
func NewApiV1() IApiV1 {
	return &ApiV1{
		IFileApi:  files.NewFilesApi(),
		IUserApi:  users.NewUsersApi(),
		IFoodsApi: foods.NewFoodApi(),
	}
}
