package controller

import (
	"TimeIimit/service"
	"github.com/gin-gonic/gin"
)

//上传单个文件
func HandleUploadFile(c *gin.Context)  {
	service.HandleUploadFile(c)
}

//上传多个文件
func HandleUploadMultipartFile(c *gin.Context)  {
	service.HandleUploadMultipartFile(c)
}

//下载文件
func HandleDownloadFile(c *gin.Context)  {
	service.HandleDownloadFile(c)
}

//注册
func Register(c *gin.Context)  {
	service.Register(c)
}

//登录
func Login(c *gin.Context) {
	service.Login(c)
}

//登出
func Logout(c *gin.Context)  {
	service.Logout(c)
}
