package middleware

import "github.com/gin-gonic/gin"

//设置是否登出检验
var set = make(map[string]string)

//登录时进行记录，登录期间阻止其他同一账号的登录
func SetSession(c *gin.Context,username string)string  {
	if set[username] == username {
		return "用户已登录"
	}
	c.SetCookie("session", username, 3600,
		"/", "localhost", 4, false, true)
	session, _ := c.Cookie("session")
	set[session]=session
	return ""
}

//验证是否登出
func CheckSession(username string)string {
	if set[username] != username {
		return "FAIL"
	}
	return "OK"
}

//登出
func OutSession(username string){
	delete(set,username)
}

//func RedisSession() (gin.HandlerFunc) {
//	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
//	return sessions.Sessions("Session",store)
//}
//
//func CheckSession(c *gin.Context)  {
//	session :=sessions.Default(c)
//	var t int64
//	ti := session.Get("time")
//	if t > time.Now().Unix(){
//
//	}
//}
//
//func SetSession(c *gin.Context,username string) {
//	session := sessions.Default(c)
//	session.Set("time",time.Now().Unix() + 60*30)
//	session.Save()
//}
