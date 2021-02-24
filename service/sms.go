package service

import (
	sms "ZhiHu/sms_Tencent"
	"github.com/gin-gonic/gin"
)

//发送带验证码的短信
func SendSms(ctx *gin.Context)(bool){
	tel := ctx.Request.FormValue("Telephone")
	test := ctx.Request.FormValue("Test")

	if test!="afhk8492"{
		return false
	}

	res:=sms.SendSms(tel)

	return res
}

//检查手机号与验证码是否一致
func CheckTelVer(ctx *gin.Context)(bool,int){

	tel := ctx.Request.FormValue("Telephone")
	verCode := ctx.Request.FormValue("VerCode")

	//手机验证码检测
	res := sms.CheckTelVer(tel,verCode)

	//用户状态检测
	UserState := sms.CheckTel(tel)

	return res,UserState
}