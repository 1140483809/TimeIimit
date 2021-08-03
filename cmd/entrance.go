package cmd

import (
	"TimeIimit/controller"
	"TimeIimit/middleware"
	"github.com/gin-gonic/gin"
)

func Entrance(){
	r := gin.Default()

	//用户路由
	user := r.Group("/user")
	{
		//注册用户
		user.GET("/register",controller.Register)
		//登录
		user.GET("/login",controller.Login)
		//登出
		user.GET("logout",controller.Logout)
	}

	//上传下载路由
	UpDown := r.Group("/updown")
	{
		// 限制表单上传大小 8MB，默认为32MB
		r.MaxMultipartMemory = 8 << 20
		//读取一个文件，保存到本地
		UpDown.POST("/upload",middleware.MidJwt(),controller.HandleUploadFile)
		//读取多个文件，保存到本地
		UpDown.POST("/upmultipart",middleware.MidJwt(),controller.HandleUploadMultipartFile)
		//下载文件
		UpDown.GET("/download",middleware.MidJwt(),controller.HandleDownloadFile)
	}

	//8080端口运行
	r.Run()
}
