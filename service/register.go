package service

import (
	. "TimeIimit/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//注册
func Register(c *gin.Context)  {
	user := new(User)
	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	message := RegisterUser(user.Username,user.Password)
	c.JSON(200,gin.H{
		"message":message,
	})
}
