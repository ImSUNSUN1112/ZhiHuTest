package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	//数据库信息
	DB_Driver = ""
)

func MysqlInit()*sql.DB{
	db,err := sql.Open("mysql",DB_Driver)
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}

	db.SetMaxOpenConns(10000)

	err = db.Ping()
	if err !=nil{
		fmt.Printf("mysql connect failed:%v", err)
	}
	DB = db
	fmt.Println("mysql connect success")
	return DB
}
