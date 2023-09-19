package foods

import (
	"cld-quicker-go/pkg"
	"github.com/gocolly/colly/v2"
)

type CommonFood struct {
	Cover string `json:"cover"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

// GetCommonFoodsLogic 获取常见的食物类型列表
func GetCommonFoodsLogic() ([]*CommonFood, error) {
	cldC := pkg.NewCLDColly(CommonFoodsList)
	CommonFoodsList := make([]*CommonFood, 0)
	cldC.C.OnHTML("td.borderBottom > a", func(e *colly.HTMLElement) {
		var CommonFoodsListItem = new(CommonFood)
		// 获取a标签的href属性
		href := e.Attr("href")
		// 打印结果
		// fmt.Printf("href: %s\n", href)
		CommonFoodsListItem.Link = href
		// 获取a标签中的img标签
		img := e.DOM.Find("img")
		if img != nil {
			// 获取img标签的src属性
			if val, exists := img.Attr("src"); exists {
				CommonFoodsListItem.Cover = val
			}
			if val, exists := img.Attr("alt"); exists {
				CommonFoodsListItem.Title = val
			}
		}
		CommonFoodsList = append(CommonFoodsList, CommonFoodsListItem)
	})
	err := cldC.Do()
	if err != nil {
		return CommonFoodsList, err
	}
	return CommonFoodsList, nil
}

// SearchFoodLogic 搜索食物
func SearchFoodLogic(foodsName string) {
	// url = fmt.Sprintf("%v%s", SearchFoods, foodsName)
	//cldC := pkg.NewCLDColly()
}
