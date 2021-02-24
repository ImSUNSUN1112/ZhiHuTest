package controller

import (
	"ZhiHu/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


//给手机发送验证码
func GetSms(ctx *gin.Context){
	res:= service.SendSms(ctx)
	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"success",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"failure",
		})
	}
}

//检查验证码与手机号是否一致
func CheckTelVer(ctx *gin.Context){

	res,UserState:=service.CheckTelVer(ctx)

	if res{
		ctx.JSON(http.StatusOK,gin.H{
			"code":10000,
			"message":"验证成功",
			"userState":UserState,
		})
	}else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "验证失败",
		})
	}
}

//--------------------------------------------------------------------------------------------------------------------//

//注册操作
func Register(ctx *gin.Context){
	res := service.Register(ctx)
	if res{
		ctx.JSON(http.StatusOK,gin.H{
			"code": 10000,
			"message":"注册成功",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"注册失败",
		})
	}
}

//登录操作
func Login(ctx *gin.Context){
	res,keyId:= service.Login(ctx)
	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"登录成功",
			"keyId":keyId,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"登录失败,用户名/手机号或密码出错",
		})
	}
}

//--------------------------------------------------------------------------------------------------------------------//

//拉取个人信息
func PeopleGet(ctx *gin.Context){
	res,name,collection:=service.PeopleGet(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"success",
			"Name":name,
			"Collection":collection,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"failure",
		})
	}
}

//更改个人信息
func PeopleEdit(ctx *gin.Context){

	res:=service.PeopleEdit(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"个人信息修改成功",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"个人信息修改失败",
		})
	}
}

//--------------------------------------------------------------------------------------------------------------------//

//获取一个问题的详细信息
func QuestionGet(ctx *gin.Context){
	res,queTitle,queBody,queAns:=service.QuestionGet(ctx)
	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"success",
			"QuestionTitle":queTitle,
			"QuestionBody":queBody,
			"QuestionAnswer":queAns,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"failure",
		})
	}
}

//发布一个问题
func QuestionPut(ctx *gin.Context){
	res:=service.QuestionPut(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"success",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"failure",
		})
	}
}

//回答一个问题
func QuestionAnswer(ctx *gin.Context){
	res:=service.QuestionAnswer(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"success",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"failure",
		})
	}
}

//--------------------------------------------------------------------------------------------------------------------//

//指定获取一篇文章
func ArticleGet(ctx *gin.Context){

	res,ArticleTitle,ArticleBody,ArticleUName:=service.ArticleGet(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"获取成功",
			"ArticleTitle":ArticleTitle,
			"ArticleBody":ArticleBody,
			"ArticleUName":ArticleUName,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"获取失败",
		})
	}

}

//发表一篇文章
func ArticlePut(ctx *gin.Context){

	res:=service.ArticlePut(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"发表文章成功",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"发表文章失败",
		})
	}
}

//随机获取一篇文章
func ArticleAny(ctx *gin.Context){

	res,ArticleTitle,ArticleBody,ArticleId:=service.ArticleAny(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"获取成功",
			"ArticleTitle":ArticleTitle,
			"ArticleBody":ArticleBody,
			"ArticleId":ArticleId,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"获取失败",
		})
	}
}

//收藏一篇文章
func ArticleCol(ctx *gin.Context){
	res := service.ArticleCol(ctx)

	if res{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"message":"收藏成功",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 20000,
			"message":"收藏失败",
		})
	}
}
