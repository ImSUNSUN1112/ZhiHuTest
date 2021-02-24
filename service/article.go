package service

import (
	"ZhiHu/AES"
	"ZhiHu/models"
	"github.com/gin-gonic/gin"
)

func ArticleGet(ctx *gin.Context)(bool,string,string,string){
	id := ctx.Request.FormValue("ArticleId")

	res,ArticleTitle,ArticleBody,ArticleUName:=models.ArticleGet(id)

	return res,ArticleTitle,ArticleBody,ArticleUName
}

func ArticlePut(ctx *gin.Context)bool{

	ArticleTitle:=ctx.PostForm("ArticleTitle")
	ArticleBody:=ctx.PostForm("ArticleBody")
	tel:=AES.AesDecrypt(ctx.PostForm("keyID"))

	res:=models.ArticlePut(tel,ArticleTitle,ArticleBody)

	return res
}

func ArticleAny(ctx *gin.Context)(bool,string,string,int){

	res,ArticleTitle,ArticleBody,ArticleId:=models.ArticleAny()

	return res,ArticleTitle,ArticleBody,ArticleId
}

func ArticleCol(ctx *gin.Context)bool{
	tel:=AES.AesDecrypt(ctx.PostForm("keyID"))
	id:=ctx.PostForm("artID")

	res:=models.ArticleCol(tel,id)

	return res
}