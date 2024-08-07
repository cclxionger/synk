package controller

import (
	"log"
	"net/http"
	"path"
	"path/filepath"

	examdir "github.com/cclxionger/synk/server/exam-dir"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FilesController(c *gin.Context) {
	// 前端给的表单是raw
	file, err := c.FormFile("raw")
	if err != nil {
		log.Fatal(err)
	}
	// 获取当前目录,如果没有就创建
	uploads := examdir.GetUpLoadsDir()
	// 文件名
	filename := uuid.New().String()
	fullpath := path.Join(uploads, filename+filepath.Ext(file.Filename))
	fileErr := c.SaveUploadedFile(file, fullpath)
	if fileErr != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})

}
