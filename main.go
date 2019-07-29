package main

import (
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	router := gin.Default()
	// 图书列表 getbooklist
	router.GET("/getBookList", getBookList)
	// 章节列表 getchapterlist/:bookid
	router.GET("/getChapterList/:bookid", getChapterList)
	// 章节详情 getchapterdetail/:bookid/:chapterid
	router.GET("/getChapterDetail/:bookid/:chapterid", getChapterDetail)

	return router
}
func getBookList(c *gin.Context) {
	c.JSON(200, gin.H{
		"iR🧒":  0,
		"sMsg": "ok",
		"data": gin.H{
			"0": gin.H{"id": "1", "name": "测试图书1"},
			"1": gin.H{"id": "2", "name": "测试图书2"},
			"2": gin.H{"id": "3", "name": "测试图书3"},
			"3": gin.H{"id": "4", "name": "测试图书4"},
			"4": gin.H{"id": "5", "name": "测试图书5"},
		},
	})
}

func getChapterList(c *gin.Context) {
	bookid := c.Param("bookid")
	c.JSON(200, gin.H{
		"iR🧒":  0,
		"sMsg": "ok",
		"data": gin.H{
			"0": gin.H{"id": "1", "b_id": bookid, "name": "测试章节1"},
			"1": gin.H{"id": "2", "b_id": bookid, "name": "测试章节2"},
			"2": gin.H{"id": "3", "b_id": bookid, "name": "测试章节3"},
			"3": gin.H{"id": "4", "b_id": bookid, "name": "测试章节4"},
			"4": gin.H{"id": "5", "b_id": bookid, "name": "测试章节5"},
		},
	})
}
func getChapterDetail(c *gin.Context) {
	bookid := c.Param("bookid")
	chapterid := c.Param("chapterid")

	c.JSON(200, gin.H{
		"iR🧒":  0,
		"sMsg": "ok",
		"data": gin.H{
			"id":      chapterid,
			"b_id":    bookid,
			"content": "图书内容在这里 巴拉巴拉巴拉巴拉巴拉",
		},
	})
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
