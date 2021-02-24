package sms_Tencent

import (
	"ZhiHu/dao"
	"fmt"
	//随机验证码获取
	"math/rand"
	"strconv"
	"time"
	//获取腾讯云sms服务
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	//锁
)

func SendSms(tel string)(bool){

	//号码检查
	if len(tel)!=11 {
		return false
	}


	credentil := common.NewCredential(
		"",
		"",
		)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client,_:=sms.NewClient(credentil,"ap-guangzhou",cpf)
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr("1400486149")
	request.Sign = common.StringPtr("SUNSUN跟张小娃子")
	request.SessionContext = common.StringPtr("test")

	//验证码获取
	rand.Seed(time.Now().Unix())
	ver:=rand.Intn(999999)
	if ver<100000{//保证获取一个六位数的验证码
		ver+=500000
	}

	//模板参数:验证码导入
	request.TemplateParamSet = common.StringPtrs([]string{strconv.Itoa(ver)})
	//模板 ID
	request.TemplateID = common.StringPtr("870637")

	//保存手机号与验证码
	addToTelTemp(tel,strconv.Itoa(ver))

	/* 下发手机号码，采用 e.164 标准，+[国家或地区码][手机号]
	 * 例如+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	request.PhoneNumberSet = common.StringPtrs([]string{"+86"+tel})

	//执行发送操作
	client.SendSms(request)

	return true
}

//将手机号与验证码加入到临时数据中
func addToTelTemp(tel string,ver string){
	if checkTelExist(tel){
		sqlStr := "update sms set vercode=? where tel = ?"
		_, err := dao.DB.Exec(sqlStr, ver, tel)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return
		}
	}else {
		sqlStr := "insert into sms(tel,vercode) values (?,?)"
		_, err := dao.DB.Exec(sqlStr, tel, ver)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}

}

//检测是否为已经发送过验证码的手机号
func checkTelExist(tel string)bool{

	var telDB string

	sqlStr := "select tel from sms where tel=?"
	err := dao.DB.QueryRow(sqlStr, tel).Scan(&telDB)
	if err != nil {
		return false
	}
	//在数据库中查找到了与目标手机号相同的手机号，返回true执行更新操作
	if tel==telDB{
		return true
	}
	//未找到相同的手机号
	return false
}

//检测该手机号是否已被注册
func CheckTel(tel string)int{
	sqlStr := "select tel from user where tel=?"
	var telDB string
	err := dao.DB.QueryRow(sqlStr, tel).Scan(&telDB)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return 0
	}
	if tel==telDB{//在数据库中检查到了相同的手机号，为已注册用户
		return 10001
	}else {
		return 10002
	}
}

//检测手机号与验证码是否相同
func CheckTelVer(tel string,ver string)bool{
	sqlStr := "select vercode from sms where tel=?"
	var verDB string
	err := dao.DB.QueryRow(sqlStr, tel).Scan(&verDB)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	if ver==verDB {
		return true
	}else {
		return false
	}
}