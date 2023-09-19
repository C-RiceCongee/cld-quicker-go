package test

import (
	"cld-quicker-go/internal/api/v1/foods"
	"cld-quicker-go/pkg"
	"fmt"
	"log"
	"testing"

	"github.com/gocolly/colly/v2"
)

// Test_Colly_HTML 标签爬取 单元测试～
func Test_Colly_HTML(t *testing.T) {
	type CommonFood struct {
		Cover string `json:"cover"`
		Title string `json:"title"`
		Link  string `json:"link"`
	}
	const url = "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/"
	cldC := pkg.NewCLDColly(url)
	var CommonFoodsList []*CommonFood

	// 每次匹配到都会执行
	cldC.C.OnHTML("td.borderBottom > a", func(e *colly.HTMLElement) {
		var CommonFoodsListItem = new(CommonFood)
		// 获取a标签的href属性
		href := e.Attr("href")
		// 打印结果
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
		log.Fatal(err)
	}
}

func TestCollySearchFoodHTML(t *testing.T) {
	url := fmt.Sprintf("%s%s&pg=%v", foods.SearchFoods, "麦当劳", 0)
	fmt.Println(url)
	cldC := pkg.NewCLDColly(url)
	cldC.C.OnHTML("td.borderBottom", func(element *colly.HTMLElement) {
		find := element.DOM.Find("a")
		// 寻找a标签
		if find != nil {
			if val, exists := find.Attr("href"); exists {
				fmt.Println(val)
			}
		}
	})
	err := cldC.Do()
	if err != nil {
		log.Fatal(err)
	}
}
