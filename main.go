package main

import (
	"ZhiHu/cmd"
	"ZhiHu/dao"
)

func main(){
	dao.MysqlInit()
	cmd.Entrance()
}