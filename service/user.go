package service

import (
	"ZhiHu/AES"
	"ZhiHu/models"
	"github.com/gin-gonic/gin"
)

//--------------------------------------------------------------------------------------------------------------------//

//注册
func Register(ctx *gin.Context)bool{

	//获取数据
	tel:=ctx.PostForm("Telephone")
	username :=ctx.PostForm("Username")
	password := ctx.PostForm("Password")

	//处理
	res:= models.Register(tel,username,password)

	return res
}

//登录
func Login(ctx *gin.Context)(bool,string){
	//获取数据
	name := ctx.PostForm("Name")
	password := ctx.PostForm("Password")

	res,keyId:= models.Login(name,password)

	return res,keyId
}

//--------------------------------------------------------------------------------------------------------------------//

//拉取个人信息
func PeopleGet(ctx *gin.Context)(bool,string,string){

	tel := AES.AesDecrypt(ctx.Request.FormValue("keyID"))

	res,name,collection:=models.PeopleGet(tel)

	return res,name,collection
}

//更改个人信息
func PeopleEdit(ctx *gin.Context)bool{

	tel := AES.AesDecrypt(ctx.PostForm("keyID"))

	name := ctx.PostForm("Name")

	password := ctx.PostForm("Password")

	res:=models.PeopleEdit(tel,name,password)

	return res
}