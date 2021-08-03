## TimeIimit

#### 路由

##### 用户路由

```go
user := r.Group("/user")
{
   //注册用户
   //{
   //    "username":"",
   //    "password":""
   //}
   user.GET("/register",controller.Register)
   //登录
   user.GET("/login",controller.Login)
   //登出
   user.GET("logout",controller.Logout)
}
```

##### 上传下载路由

```go
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
```

#### 控制器

```go
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
```

#### 上传下载

##### 上传

储存在本地项目（gofile/用户名）下

http://localhost:8080/updown/upload          上传单个文件

http://localhost:8080/updown/upmultipart   上传多个文件

##### 下载

http://localhost:8080/updown/download?filename=文件名   下载此文件

#### 大致实现的功能

- 注册 http://localhost:8080/user/register
- 登录 http://localhost:8080/user/login 
- 注销 http://localhost:8080/user/logout
- 上传文件
- 下载文件
- 登录期间阻止其他同一账号登录请求