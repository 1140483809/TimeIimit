package service

import (
	"TimeIimit/middleware"
	. "TimeIimit/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//用户登录
func Login(c *gin.Context)  {
	user := new(User)
	if err := c.BindJSON(user); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	message := CheckLogin(user.Username,user.Password)
	var token string
	if message == "登录成功"{
		token , _ = middleware.MidNewJwt(user.Username)
		ok := middleware.SetSession(c,user.Username)
		if ok=="用户已登录" {
			c.JSON(200,gin.H{
				"message":ok,
			})
			return
		}
	}
	c.JSON(200,gin.H{
		"message":message,
		"token":token,
	})
}
