package models

import (
	"ZhiHu/dao"
	"math/rand"
	"time"
)

type art struct {
	id int `json:"id"`
	tel string `json:"tel"`
	title string `json:"title"`
	body string `json:"body"`
	hot int `json:"hot"`
}

//获取指定的一篇文章
func ArticleGet(id string)(bool,string,string,string){

	sqlStr := "select tel,title,body from art where id=?"
	var artDB art
	err := dao.DB.QueryRow(sqlStr,id).Scan(&artDB.tel,&artDB.title,&artDB.body)
	if err != nil{
		return false,"","",""
	}

	sqlStr = "select name from user where tel=?"
	var nameDB string
	err = dao.DB.QueryRow(sqlStr,artDB.tel).Scan(&nameDB)
	if err != nil{
		return false,"","",""
	}

	//文章热度增加
	//获取当前热度
	sqlStr = "select hot from art where id=?"
	var hotDB int
	err = dao.DB.QueryRow(sqlStr,id).Scan(&hotDB)
	if err != nil{
		return false,"","",""
	}
	hotDB+=1
	sqlStr = "update art set hot =? where id = ?"
	_,err=dao.DB.Exec(sqlStr,hotDB,id)
	if err != nil{
		return false,"","",""
	}

	return true,artDB.title,artDB.body,nameDB
}

//发布一篇文章
func ArticlePut(tel string,title string,body string)bool{
	sqlStr := "insert into art(tel,title,body ) values (?,?,?)"
	_, err := dao.DB.Exec(sqlStr,tel,title ,body)
	if err != nil {
		return false
	}
	return true
}

//随机获取一篇文章
func ArticleAny()(bool,string,string,int){

	rand.Seed(time.Now().Unix())
	id:=rand.Intn(20)

	sqlStr := "select id,title,body from art where id=?"
	var artDB art
	err := dao.DB.QueryRow(sqlStr,id).Scan(&artDB.id,&artDB.title,&artDB.body)
	if err != nil{
		return false,"","",0
	}
	return true,artDB.title,artDB.body,artDB.id
}

//收藏一篇文章
func ArticleCol(tel string,id string)bool{

	sqlStr := "select collectionArt from user where tel=?"

	var idDB string

	err := dao.DB.QueryRow(sqlStr,tel).Scan(&idDB)
	if err != nil{
		return false
	}

	idDB=idDB+" "+id

	sqlStr = "update user set collectionArt =? where tel = ?"
	_,err=dao.DB.Exec(sqlStr,idDB,tel)
	if err != nil{
		return false
	}

	return true
}
