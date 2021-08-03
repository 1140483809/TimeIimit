package main

import (
	"TimeIimit/cmd"
	"TimeIimit/models"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	models.SetupEngine()  //连接数据库
	cmd.Entrance()        //运行
}
