package module

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	LastChapter string `json:"last_chapter"`
	Finished    int    `json:"finished"`
}
type Chapter struct {
	Id    int    `json:"id"`
	Bid   int    `json:"b_id"`
	Title string `json:"title"`
}
type ChapterContent struct {
	Id      int    `json:"id"`
	Bid     int    `json:"b_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	PreId   int    `json:"pre_id"`
	AfterId int    `json:"after_id"`
}

//获取图书列表
func GetBooks(c *gin.Context) {
	rows, err := DBH.Query("select `id`,`name`,`last_chapter`,`finished` from book ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var books []Book

	for rows.Next() {
		var id int
		var name string
		var lastChapter string
		var finished int
		err := rows.Scan(&id, &name, &lastChapter, &finished)
		if err != nil {
			panic(err)
		}
		books = append(books, Book{
			id,
			name,
			lastChapter,
			finished,
		})
	}
	Responds(0, "ok", books, c)
}

//获取图书的章节列表
func GetBookChapters(c *gin.Context) {
	bId, _ := strconv.Atoi(c.Param("b_id"))
	stmtOut, err := DBH.Prepare("select `id`,`title` from book_chapter where b_id=? ")
	if err != nil {
		panic(err)
	}
	defer stmtOut.Close()
	rows, err := stmtOut.Query(bId)
	if err != nil {
		panic(err)
	}
	var chapters []Chapter

	for rows.Next() {
		var id int
		var title string
		err := rows.Scan(&id, &title)
		if err != nil {
			panic(err)
		}
		chapters = append(chapters, Chapter{
			id,
			bId,
			title,
		})
	}
	Responds(0, "ok", chapters, c)
}

func GetChapterContent(c *gin.Context) {
	bId, _ := strconv.Atoi(c.Param("b_id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var title string
	var content string
	stmtOut, err := DBH.Prepare("SELECT `title`,`content` FROM book_chapter WHERE b_id=? and id= ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(bId, id).Scan(&title, &content) // WHERE number = 1

	if (err != nil) && (err != sql.ErrNoRows) {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	//PreId   int    `json:"pre_id"`
	//	AfterId int    `json:"after_id"`

	var PreId int
	var AfterId int
	err = DBH.QueryRow("select id from book_chapter where b_id = ? and id < ? order by id desc limit 1", bId, id).Scan(&PreId)
	if (err != nil) && (err != sql.ErrNoRows) {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	err = DBH.QueryRow("select id from book_chapter where b_id = ? and id > ? order by id asc limit 1", bId, id).Scan(&AfterId)
	if (err != nil) && (err != sql.ErrNoRows) {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	chapterDetail := ChapterContent{
		id,
		bId,
		title,
		content,
		PreId,
		AfterId,
	}
	Responds(0, "ok", chapterDetail, c)
}

//https://github.com/go-sql-driver/mysql/wiki/Examples
