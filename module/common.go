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
	Passwd string `ini:"passwd"`
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

func Login(c *gin.Context) {
	//user := c.Param("name")
	//passwd := c.Param("name")
	Responds(0, "logined", "", c)

}

func LoginOut(c *gin.Context) {
	Responds(0, "LoginOut", "", c)
}
