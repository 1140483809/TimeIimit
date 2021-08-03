package service

import (
	"github.com/gin-gonic/gin"
	"path"
)

//各种文件
var HttpContentType = map[string]string{
	".avi":  "video/avi",
	".mp3":  "audio/mp3",
	".mp4":  "video/mp4",
	".wmv":  "video/x-ms-wmv",
	".asf":  "video/x-ms-asf",
	".rm":   "application/vnd.rn-realmedia",
	".rmvb": "application/vnd.rn-realmedia-vbr",
	".mov":  "video/quicktime",
	".m4v":  "video/mp4",
	".flv":  "video/x-flv",
	".jpg":  "image/jpeg",
	".png":  "image/png",
	".txt":  "text/plain",
}

//下载文件
func HandleDownloadFile(c *gin.Context) {
	filename := c.Query("filename")
	username := c.GetString("name")
	filePath := path.Join("./gofile/",username)
	filePath = path.Join(filePath,filename)
	//获取文件名称带后缀
	fileNameWithSuffix := path.Base(filePath)
	//获取文件的后缀
	fileType := path.Ext(fileNameWithSuffix)
	//获取文件类型对应的http ContentType 类型
	fileContentType := HttpContentType[fileType]
	c.Header("Content-Type", fileContentType)
	c.File(filePath)
}