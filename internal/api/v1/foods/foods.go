package foods

import (
	"cld-quicker-go/internal/api"
	"github.com/gin-gonic/gin"
)

type IFoodsApi interface {
	api.IBase
	GetCommonFoods() func(ctx *gin.Context)
}

type FoodApi struct {
}

func (f FoodApi) SetPrefix() string {
	return "/api/v1/foods"
}

func (f FoodApi) GetCommonFoods() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
	}
}

func NewFoodApi() IFoodsApi {
	return &FoodApi{}
}
