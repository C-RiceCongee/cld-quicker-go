package test

import (
	"cld-quicker-go/pkg"
	"fmt"
	"testing"
)

func Test_Colly_HTML(t *testing.T) {
	// 标签爬取 单元测试～
	// 主页
	const url = "https://www.fatsecret.cn/%E7%83%AD%E9%87%8F%E8%90%A5%E5%85%BB/"
	colly := pkg.NewColly(url)
	//fmt.Println(111)
	colly.OnHTML("h2", "")
	err := colly.DoVisit()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range colly.HTML {
		fmt.Println(i.Text)
		if i == nil {
			break
		}
	}
}
