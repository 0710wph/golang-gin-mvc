package main

import (
	"gin/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//database.InitMysql()
	router := routers.InitRouter()
	router.Run(":8181")
}