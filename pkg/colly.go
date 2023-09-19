package pkg

import "C"
import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type CLDColly struct {
	url  string
	c    *colly.Collector
	HTML chan *colly.HTMLElement
}

func (cld *CLDColly) DoVisit() error {
	return cld.c.Visit(cld.url)
}

func (cld *CLDColly) OnHTML(tag string, attr string) {
	cld.c.OnHTML(tag, func(e *colly.HTMLElement) {
		cld.HTML <- e
	})
}

func (cld *CLDColly) OnRequest() {
	cld.c.OnRequest(func(r *colly.Request) {
		fmt.Print(r.URL)
	})
}

func NewColly(url string) *CLDColly {
	return &CLDColly{
		url:  url,
		c:    colly.NewCollector(),
		HTML: make(chan *colly.HTMLElement, 10),
	}
}
