package module

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection string `ini:"connection"`
	Host       string `ini:"host"`
	Port       string `ini:"port"`
	Database   string `ini:"database"`
	Username   string `ini:"username"`
	Password   string `ini:"password"`
}
type Redis struct {
	Host     string `ini:"host"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
}
type Server struct {
	Port   string `ini:"port"`
	Token  string `ini:"token"`
	Admin  string `ini:"admin"`
	Pass   string `ini:"pass"`
	Domain string `ini:"domain"`
}
type AppConf struct {
	Database `ini:"database"`
	Redis    `ini:"redis"`
	Server   `ini:"server"`
}

var DBH *sql.DB
var AppConfig AppConf

func init() {
	err := ini.MapTo(&AppConfig, ".env.ini")
	fmt.Println(AppConfig)
	if err != nil {
		panic(err)
	}
	dsn := AppConfig.Database.Username + ":" + AppConfig.Database.Password + "@tcp(" + AppConfig.Database.Host + ":"
	dsn += AppConfig.Database.Port + ")/" + AppConfig.Database.Database + "?charset=utf8"

	DBH, err = sql.Open(AppConfig.Database.Connection, dsn)
	if err != nil {
		panic(err)
	}
	err = DBH.Ping()
	if err != nil {
		panic(err)
	}
}
func Responds(iRet int, sMsg string, data interface{}, c *gin.Context) {
	responseData := make(map[string]interface{})
	responseData["iRet"] = iRet
	responseData["sMsg"] = sMsg
	responseData["data"] = data
	c.JSON(0, responseData)
	c.Abort() // 结束当前请求
}

func IsLogin(c *gin.Context) {
	val, _ := c.Cookie("login")
	if val == "1" {
		Responds(0, "login", "1", c)
	} else {
		Responds(0, "not login", "-1", c)
	}
}

func Login(c *gin.Context) {
	user := c.Request.FormValue("name")
	pass := c.Request.FormValue("pass")

	if user == AppConfig.Server.Admin && pass == AppConfig.Server.Pass {
		c.SetCookie("login", "1", 86400, "/", AppConfig.Server.Domain, false, true)
		Responds(0, "login", "", c)
	} else {
		Responds(-1, "wrong login info", "", c)
	}
}

func LoginOut(c *gin.Context) {
	c.SetCookie("login", "1", -1, "/", AppConfig.Server.Domain, false, true)
	Responds(0, "LoginOut", "", c)
}
