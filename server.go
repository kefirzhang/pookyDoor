package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pookyDoor/module"
)

var API_TOKEN string

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")
		if token == "" {
			module.Responds(401, "API token required", "", c)
			return
		}
		if token != module.AppConfig.Server.Token {
			module.Responds(401, "Invalid API token", "", c)
			return
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/Login", module.Login)
	router.GET("/LoginOut", module.LoginOut)
	router.Use(TokenAuth())
	router.GET("/GetBooks", module.GetBooks)                             // 图书列表
	router.GET("/GetBookChapters/:b_id", module.GetBookChapters)         // 章节列表
	router.GET("/GetChapterContent/:b_id/:id", module.GetChapterContent) // 章节详情
	return router
}
func main() {
	r := setupRouter()
	fmt.Println(module.AppConfig)
	r.Run(":" + module.AppConfig.Server.Port)
}
