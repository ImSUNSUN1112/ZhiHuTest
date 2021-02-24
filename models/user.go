package models

import (
	"ZhiHu/dao"
	"ZhiHu/AES"
	"fmt"
)

type user struct {
	KeyId string `json:"keyId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Collection string `json:"collection"`
}

//注册
func Register(tel string,name string,password string,)(bool){

	//检查该用户是否已经存在
	if checkUserTelExist(tel){
		return false
	}
	//预处理
	stmt,err := dao.DB.Prepare("insert into user (tel,name,password)values (?,?,?)")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	//关闭连接
	defer stmt.Close()
	//插入
	_, err = stmt.Exec(tel,name,password)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	//注册成功
	return true
}

//查询用户手机号是否已经存在
func checkUserTelExist(tel string)bool{
	stmt, _ := dao.DB.Prepare("select tel from user where tel = ?")
	defer stmt.Close()
	//手机号作为唯一凭证进行筛查
	rows, _ := stmt.Query(tel)
	defer rows.Close()
	//读取结果
	for rows.Next() {
		//获取数据(如果存在
		var telDB string
		rows.Scan(&telDB)

		if tel==telDB{//找到了相同的注册用户
			return true
		}
	}
	//没有在数据库中找到相同的注册用户
	return false
}

//登录
func Login(name string,password string)(bool,string){

	stmt, err := dao.DB.Prepare("select password from user where tel = ?")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v", err)
		return false,""
	}

	defer stmt.Close()

	rows, err := stmt.Query(name)

	if err!=nil{
		return false,""
	}

	defer rows.Close()
	//读取结果
	for rows.Next() {
		//获取数据库中的密码
		var passwordDB string
		rows.Scan(&passwordDB)
		//密码对比
		if password==passwordDB{//找到了相同的注册用户
			return true,AES.AesEncrypt(name)
		}
	}

	//没有在数据库中找到相同的注册用户
	return false,""
}

//更改个人信息
func PeopleEdit(tel string,name string,password string)bool{

	sqlStr := "update user set name =? where tel = ?"
	_, err := dao.DB.Exec(sqlStr, name, tel)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return false
	}
	sqlStr = "update user set password=? where tel = ?"
	_, err = dao.DB.Exec(sqlStr, password, tel)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return false
	}

	return true
}

//拉取个人信息
func PeopleGet(tel string)(bool,string,string){
	sqlStr := "select name,collection from user where tel=?"
	var u user
	err := dao.DB.QueryRow(sqlStr, tel).Scan(&u.Username,&u.Collection)
	if err != nil {
		return false,"",""
	}
	return true,u.Username,u.Collection
}




