package main

import (
	"github.com/gin-gonic/gin"
	"pookyDoor/module"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		val, _ := c.Cookie("login")
		if val != "1" {
			module.Responds(-1, "not login", "-1", c)
			return
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/Login", module.Login)
	router.GET("/LoginOut", module.LoginOut)
	router.GET("/IsLogin", module.IsLogin)
	//下面操作需要登陆
	router.Use(CheckLogin())
	router.GET("/GetBooks", module.GetBooks)                             // 图书列表
	router.GET("/GetBookChapters/:b_id", module.GetBookChapters)         // 章节列表
	router.GET("/GetChapterContent/:b_id/:id", module.GetChapterContent) // 章节详情
	return router
}
func main() {
	r := setupRouter()
	r.Run(":" + module.AppConfig.Server.Port)
}
