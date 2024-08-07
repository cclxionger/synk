package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	examdir "github.com/cclxionger/synk/server/exam-dir"
	"github.com/gin-gonic/gin"
)

func TextsController(c *gin.Context) {
	// 上传文字的类型是 raw:....
	type Json struct {
		Raw string
	}
	var json Json
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	} else {
		// 把这个json数据写入到一个数据库里面，然后从数据库里面取
		// 现在不用数据库，而是创建一个文件，把这个文件交给前端

		// 获取当前目录,如果没有就创建
		examdir.GetUpLoadsDir()
		// 文件名是当天的日期，注意文件名要符合文件名规范(比如不能有:)
		fileName := time.Now().Format("2006-01-02_15-04-05")
		// 文件名如果需要随机生成，则是这个
		// fileName := uuid.New().String()
		fullFilePath := filepath.Join("uploads", fileName+".txt")
		// 创建文件
		file, err := os.OpenFile(fullFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		_, err = file.WriteString(json.Raw)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"url": "/" + fullFilePath,
		})

	}

}
