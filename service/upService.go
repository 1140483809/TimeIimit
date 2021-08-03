package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

// HandleUploadFile 上传单个文件
func HandleUploadFile(c *gin.Context) {
	username := c.GetString("name")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "文件上传失败",
		})
		return
	}
	fmt.Println(username)
	dst := path.Join("./gofile/",username)
	dst1 := path.Join(dst,"/",file.Filename)
	fmt.Println(dst)
	a := 0
save:if err = c.SaveUploadedFile(file,dst1);
	err != nil && a==0 {
		err := os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return
		}
		a++
		goto save
	}
	if err != nil {
		c.JSON(http.StatusBadGateway,gin.H{
			"response":"文件保存失败",
		})
		return
	}
	fmt.Println(dst1)
	c.JSON(http.StatusOK, gin.H{
		"response": "上传成功",
	})
}
//上传多个文件
func HandleUploadMultipartFile(c *gin.Context)  {
	username := c.GetString("name")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadGateway,gin.H{
			"response":"上传失败",
		})
	}
	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		dst := path.Join("./gofile/",username)
		dst1 := path.Join(dst,"/",file.Filename)
		a := 0
		save:if err := c.SaveUploadedFile(file, dst1)
			err != nil && a==0 {
			err := os.MkdirAll(dst, os.ModePerm)
			if err != nil {
				return
			}
			a++
			goto save
		}
		if err != nil {
			c.JSON(http.StatusBadGateway,gin.H{
				"response":"文件保存失败",
			})
			return
		}
		fmt.Println(dst1)
	}
	c.JSON(http.StatusOK,gin.H{
		"response":"文件上传成功",
	})
}