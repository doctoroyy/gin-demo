package spider

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Spider struct {
	c *colly.Collector
}

func NewSpider() *Spider {
	return &Spider{
		c: colly.NewCollector(colly.DetectCharset(), colly.AllowURLRevisit()),
	}
}

func (s *Spider) GetChapterUrls(chapterQuery string, url string) interface{} {
	chapterUrls := []string{}

	s.c.OnHTML(chapterQuery, func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			href := e.Attr("href")
			chapterUrls = append(chapterUrls, href)
		})
	})

	s.c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error:", err.Error())
	})

	s.c.Visit(url)

	s.c.Wait()

	return chapterUrls
}

func (s *Spider) GetContent(query string, url string) interface{} {
	content := ""

	s.c.OnHTML(query, func(e *colly.HTMLElement) {
		content = e.Text
	})

	s.c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error:", err.Error())
	})

	s.c.Visit(url)

	s.c.Wait()

	return content
}
