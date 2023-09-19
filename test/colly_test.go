package test

import (
	"cld-quicker-go/internal/api/v1/foods"
	"cld-quicker-go/pkg"
	"fmt"
	"log"
	"net/url"
	"strings"
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
	const aimURL = "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/"
	cldC := pkg.NewCLDColly(aimURL)
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
func TestChineseToUTF(t *testing.T) {
	str := "酱香饼"
	values := url.Values{}
	values.Add("q", str)
	encode := values.Encode()
	fmt.Println(encode)
}
func TestCollySearchFoodHTML(t *testing.T) {
	aimURL := fmt.Sprintf("%s%s", foods.SearchFoods, pkg.URIEncoder([]string{"q",
		"焦糖瓜子", "pg", "0"}))
	cldC := pkg.NewCLDColly(aimURL)
	cldC.C.OnHTML("td.borderBottom", func(element *colly.HTMLElement) {
		aProminent := element.DOM.Find("a.prominent")
		// 寻找a标签
		if aProminent != nil {
			fmt.Println(aProminent.Text())
			if val, exists := aProminent.Attr("href"); exists {
				fmt.Println(val)
			}
		}
		divSmallText := element.DOM.Find("div.smallText")
		if divSmallText != nil {
			smallText := strings.TrimSpace(divSmallText.Text())
			smallTextArr := strings.Split(smallText, "\n")
			smallTextArrNew := make([]string, 0)
			for i := 0; i < len(smallTextArr); i++ {
				if len(strings.TrimSpace(smallTextArr[i])) > 0 {
					smallTextArrNew = append(smallTextArrNew, strings.TrimSpace(smallTextArr[i]))
				}
			}
			fmt.Println(smallTextArrNew[0])
		}
	})
	err := cldC.Do()
	if err != nil {
		log.Fatal(err)
	}
}
