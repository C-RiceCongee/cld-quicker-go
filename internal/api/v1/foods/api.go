package foods

import (
	"cld-quicker-go/com"
	"cld-quicker-go/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IFoodsApi interface {
	api.IBase
	GetCommonFoods() func(ctx *gin.Context)
	SearchFood() func(ctx *gin.Context)
}

type FoodApi struct {
}

func (f FoodApi) SetPrefix() string {
	return "/api/v1/foods"
}

// GetCommonFoods 获取常见的食物类型列表
func (f FoodApi) GetCommonFoods() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cf, err := GetCommonFoodsLogic()
		if err != nil {
			com.ResponseData(ctx, http.StatusOK, cf, "获取食物类型失败")
			return
		}
		com.ResponseData(ctx, com.CodeSuccess, cf)
	}
}

// SearchFood 查询食物
func (f FoodApi) SearchFood() func(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
func NewFoodApi() IFoodsApi {
	return &FoodApi{}
}
