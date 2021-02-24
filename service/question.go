package service

import (
	"ZhiHu/AES"
	"ZhiHu/models"
	"github.com/gin-gonic/gin"
)


func QuestionGet(ctx *gin.Context)(bool,string,string,string){
	//问题id获取
	queId:= ctx.PostForm("QuestionId")

	res,queTitle,queBody,queAns:=models.QuestionGet(queId)

	return res,queTitle,queBody,queAns
}

func QuestionPut(ctx *gin.Context)bool{

	//获取问题标题
	queTitle:=ctx.PostForm("QuestionTitle")
	//获取问题主体
	queBody:=ctx.PostForm("QuestionBody")
	//获取提问人信息
	tel:=AES.AesDecrypt(ctx.PostForm("QuestionkeyId"))
	//执行
	res:=models.QuestionPut(queTitle,queBody,tel)

	//返回执行情况
	return res
}

func QuestionAnswer(ctx *gin.Context)bool{
	//获取回答主体
	ansBody:=ctx.PostForm("AnswerBody")
	//获取回答人
	ansName:=ctx.PostForm("AnswerName")
	//获取问题Id
	id:=ctx.PostForm("QuestionId")

	res:=models.QuestionAnswer(ansBody,ansName,id)

	return res
}