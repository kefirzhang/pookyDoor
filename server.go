package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
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
		if token != API_TOKEN {
			module.Responds(401, "Invalid API token", "", c)
			return
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(TokenAuth())
	router.GET("/GetBooks", module.GetBooks)                             // 图书列表
	router.GET("/GetBookChapters/:b_id", module.GetBookChapters)         // 章节列表
	router.GET("/GetChapterContent/:b_id/:id", module.GetChapterContent) // 章节详情
	return router
}
func main() {
	cfg, err := ini.Load(".env.ini")
	if err != nil {
		panic(err)
	}
	API_TOKEN = cfg.Section("server").Key("server_api").String()
	module.Setup() // 初始化
	r := setupRouter()
	r.Run(":" + cfg.Section("server").Key("server_port").String())
}
