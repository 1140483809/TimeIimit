package models

import (
	"github.com/go-xorm/xorm"
	"log"
)

//用户账号密码
type User struct {
	Username string `xorm:"pk" json:"username"`
	Password string `json:"password"`
}

var Usermap = make(map[string]User)
var EngineSql *xorm.Engine

//连接数据库，并将账号密码提出
func SetupEngine()*xorm.Engine  {
	var conn = "root:Yy13883129603@/sqlsql?charset=utf8"
	EngineSql,err := xorm.NewEngine("mysql",conn)
	if err != nil {
		panic("mysql connect fail")
	}
	err = EngineSql.Sync2(new(User))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = EngineSql.Find(Usermap)
	if err != nil {
		log.Fatal(err.Error())
	}
	return EngineSql
}
