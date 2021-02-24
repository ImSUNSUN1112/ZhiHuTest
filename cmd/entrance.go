package cmd

import (
	"ZhiHu/controller"
	"ZhiHu/middlewares"
	"github.com/gin-gonic/gin"
)

func Entrance(){

	//启动路由
	router := gin.Default()

	//添加中间件解决跨域问题
	router.Use(middlewares.Cors())

	//托管
	routerGroup := router.Group("ZhiHu")
	{
		//登录注册
		//获取短信验证码
		routerGroup.GET("/getSms",controller.GetSms)
		//验证验证码
		routerGroup.GET("/checkSms",controller.CheckTelVer)
		//注册
		routerGroup.POST("/register", controller.Register)
		//登录
		routerGroup.POST("/login", controller.Login)

		//个人信息
		//拉取一个人的个人信息
		routerGroup.GET("/people",controller.PeopleGet)
		//更改一个人的个人信息
		routerGroup.POST("/people/edit",controller.PeopleEdit)


		//问题相关
		//获取一个问题的所有信息,检索信息:问题ID
		routerGroup.GET("/question",controller.QuestionGet)
		//发布一个问题
		routerGroup.POST("/question/put",controller.QuestionPut)
		//回答一个问题
		routerGroup.POST("/question/answer",controller.QuestionAnswer)

		//文章相关
		//获取一个文章的所有信息,检索信息:文章ID
		routerGroup.GET("/article",controller.ArticleGet)
		//发布一篇文章
		routerGroup.POST("/article/put",controller.ArticlePut)
		//随机获取一篇文章用于填充主界面,返回 标题,内容,id
		routerGroup.GET("/article/any",controller.ArticleAny)
		//收藏一篇文章
		routerGroup.POST("/article/collection",controller.ArticleCol)
	}

	router.Run(":8000")
}
