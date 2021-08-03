package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//客户端信息
type MyCLaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//密钥
var MysigningKey = []byte("timelimit")


//签发jwt
func MidNewJwt(username string)(string,string) {
	cok := MyCLaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,     //生效时间一分钟前
			ExpiresAt: time.Now().Unix() + 60*30,  //到期时间半小时
			Issuer:    "lin",
		},
	}
	//哈希256加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,cok)
	s,err := token.SignedString(MysigningKey);
	if err!=nil {
		return "","FAIL"
	}
	fmt.Println(s)
	return s,"OK"
}


//检测jwt
func MidCheckJwt(token string)(*MyCLaims, string){
	//解密
	settoken,err := jwt.ParseWithClaims(token,&MyCLaims{}, func(token *jwt.Token) (interface{}, error) {
		return MysigningKey, nil
	})
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	if key, _ := settoken.Claims.(*MyCLaims); settoken.Valid{
		return key,"OK"
	}else {
		return nil,"FAIL"
	}
}

//中间件
func MidJwt() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenherder := c.Request.Header.Get("Authorization")
		code := "OK"
		if tokenherder == "" {
			code = "token不存在"
			c.JSON(http.StatusOK,gin.H{
				"message" : code,
			})
			c.Abort()
			return
		}
		checktoken := strings.SplitN(tokenherder," ",2)
		if len(checktoken) != 2 && checktoken[0] != "Bearer" {
			code = "token格式错误"
			c.JSON(http.StatusOK,gin.H{
				"message" : code,
			})
			c.Abort()
			return
		}
		key , tcode := MidCheckJwt(checktoken[1])
		if tcode == "FAIL" {
			code = "token不正确"
			c.JSON(http.StatusOK,gin.H{
				"message" : code,
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt{
			code = "token已过期"
			c.JSON(http.StatusOK,gin.H{
				"message" : code,
			})
			c.Abort()
			return
		}
		ok := CheckSession(key.Username);
		if ok == "FAIL" {
			code = "未登录"
			c.JSON(http.StatusOK,gin.H{
				"message" : code,
			})
			c.Abort()
			return
		}
		c.Set("name",key.Username)
		c.Next()
	}
}
