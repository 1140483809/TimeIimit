package service

import (
	"TimeIimit/middleware"
	"github.com/gin-gonic/gin"
	"strings"
)

//登出
func Logout(c *gin.Context)  {
	tokenherder := c.Request.Header.Get("Authorization")
	checktoken := strings.SplitN(tokenherder," ",2)
	key , _ := middleware.MidCheckJwt(checktoken[1])
	middleware.OutSession(key.Username)
	c.JSON(200,gin.H{
		"message" : "已注销",
	})
}
