package models

import (
	"ZhiHu/dao"
)

type que struct {
	id int
	tel string
	title string
	body string
	answer string
	hot int
}

//获取一个问题的详情页
func QuestionGet(Id string)(bool,string,string,string){
	sqlStr := "select title,body from que where id=?"
	var queDB que
	err := dao.DB.QueryRow(sqlStr,Id).Scan(&queDB.title,&queDB.body)
	if err != nil{
		return false,"","",""
	}

	return true,queDB.title,queDB.body,""
}

func QuestionPut(title string,body string,tel string)bool{
	sqlStr := "insert into que(tel,title,body) values (?,?,?)"
	_,err := dao.DB.Exec(sqlStr,tel,title,body)
	if err != nil {
		return false
	}
	return true
}

func QuestionAnswer(ansBody string,ansName string,id string)bool{
	answer:=ansName+":"+ansBody

	sqlStr := "select answer from que where id=?"
	var ansDB string
	err := dao.DB.QueryRow(sqlStr,id).Scan(&ansDB)
	if err != nil{
		return false
	}
	ansDB=ansDB+answer
	sqlStr = "update que set answer =? where id = ?"
	_,err=dao.DB.Exec(sqlStr,ansDB,id)
	if err != nil{
		return false
	}

	return true
}
