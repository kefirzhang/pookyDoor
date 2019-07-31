package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"pookyDoor/module"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
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
	module.Setup() // 初始化
	r := setupRouter()
	r.Run(":" + cfg.Section("server").Key("server_port").String())
}
