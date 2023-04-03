package main

import (
	"gin-demo/spider"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	QuerySelector string `json:"querySelector"`
	Url           string `json:"url"`
}

func main() {

	r := gin.Default()

	spider := spider.NewSpider()

	handleRequest := func(c *gin.Context, f func(string, string) interface{}) {
		var reqBody RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "参数错误",
				"data": nil,
			})
			return
		}

		content := f(reqBody.QuerySelector, reqBody.Url)

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": content,
		})
	}

	r.POST("/chapter", func(c *gin.Context) {
		handleRequest(c, spider.GetChapterUrls)
	})

	r.POST("/content", func(c *gin.Context) {
		handleRequest(c, spider.GetContent)
	})

	r.Run(":8090")
}
