package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pookyDoor/module"
	"strings"
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

func Cors() gin.HandlerFunc { //TODO 后续这个函数需要优化
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//下面的都是乱添加的-_-~
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin")) //TODO 安全优化点
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	router.POST("/Login", module.Login) //登陆
	//router.GET("/Login", module.Login)      //登陆
	router.GET("/LoginOut", module.LoginOut) //登出
	router.GET("/IsLogin", module.IsLogin)   //是否登陆
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
