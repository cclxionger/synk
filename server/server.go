package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cclxionger/synk/server/controller"
	websocket "github.com/cclxionger/synk/server/web-socket"
	"github.com/gin-gonic/gin"
)

var GinErrChan = make(chan error, 1)

func Run() {
	gin.SetMode("debug")
	r := gin.Default()

	// 服务静态文件
	r.StaticFS("/static", http.Dir("./server/frontend/dist")) // 访问/static/index.html就可以访问
	// 希望接收到/static 都返回上面那个html文件，不是以/static开头的，返回404
	r.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if ok := strings.HasPrefix(path, "/static"); ok {
			dirpath := "./server/frontend/dist/index.html"
			data, err := os.ReadFile(dirpath)
			if err != nil {
				GinErrChan <- err
				log.Fatalln(err)
			}
			ctx.Data(http.StatusOK, "text/html; charset=uft-8", data)

		} else {
			ctx.Status(http.StatusNotFound)
		}

	})
	r.POST("/api/v1/texts", controller.TextsController)
	r.POST("/api/v1/files", controller.FilesController)
	r.GET("/api/v1/addresses", controller.AddressesController)
	r.GET("/uploads/:path", controller.UploadsController)
	r.GET("/api/v1/qrcodes", controller.QrcodesController)
	r.GET("/ws", func(c *gin.Context) {
		websocket.HttpController(c, &websocket.Hub{})
	})
	r.Run(":27149")
}
