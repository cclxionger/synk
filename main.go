package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/cclxionger/synk/server"
	examdir "github.com/cclxionger/synk/server/exam-dir"
	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
)

// 哪里执行synk.exe文件，文件就上传到synk.exe的当前目录的/uploads下面
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

		// 获取当前目录
		fileLoads := examdir.GetUpLoadsDir()
		// 文件名是当天的日期，注意文件名要符合文件名规范(比如不能有:)
		fileName := time.Now().Format("2006-01-02_15-04-05")
		// 文件名如果需要随机生成，则是这个
		// fileName := uuid.New().String()
		fullFilePath := filepath.Join(fileLoads, fileName+".txt")
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

func main() {

	go server.Run()

	CreateUI()

}
func CreateUI() {
	ui, err := lorca.New("http://127.0.0.1:27149/static/index.html", "", 1500, 900, "--remote-allow-origins=*")
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭的实话移除所有临时文件
	defer ui.Close()
	// 想要实现 ctrl+c interpret 中断退出
	signalChan := make(chan os.Signal, 1)
	// Interrupt Signal = syscall.SIGINT 中断
	signal.Notify(signalChan, os.Interrupt)
	select {
	// 阻塞，直到有管道有信号
	case err = <-server.GinErrChan:
		log.Println(err)
	case <-signalChan:
		log.Println("interpret")
	case <-ui.Done():
		// 从任务管理器结束是中止，直接点❌，都是done
		log.Println("done")
	}
}
