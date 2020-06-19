package utils

import (
	"database/sql"
	"fmt"
)

func InitMysql() {
	fmt.Println("InitMysql....")
	if db == nil {
		db, _ = sql.Open("mysql", "weiyuan:Weiyuan!12@cnwy04@tcp(rm-bp166ow4x550532ba7o.mysql.rds.aliyuncs.com:3306)/fcweb?parseTime=true")
	}
}

