package controller

import (
	"net/http"
	"path/filepath"

	examdir "github.com/cclxionger/synk/server/exam-dir"
	"github.com/gin-gonic/gin"
)

// 把网络路径:path变成本地绝对路径，读取文件，写到http响应里面
func UploadsController(c *gin.Context) {
	if path := c.Param("path"); path != "" {
		target := filepath.Join(examdir.GetUpLoadsDir(), path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream")
		// 是去下载
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
