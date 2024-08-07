package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// 生成二维码，可以后端，也可以前端
// QrcodesController 用于处理生成二维码的HTTP请求
// 它根据URL查询参数中的"content"来生成对应的二维码图片，并以PNG格式返回给客户端
// 如果"content"参数为空，则返回400 Bad Request状态码
func QrcodesController(c *gin.Context) {
	// 从URL的查询参数中获取"content"字段的值
	if content := c.Query("content"); content != "" {
		// 使用qrcode包（假设已经引入并可用）的Encode函数生成二维码
		// Encode函数的参数包括要编码的内容、二维码的级别（此处为Medium）、以及二维码的大小（像素）
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		// 检查是否有错误发生
		if err != nil {
			// 如果有错误，记录错误信息并终止程序（在实际应用中，可能更希望返回错误信息给客户端而不是终止程序）
			log.Fatalln(err)
		}
		// 设置HTTP响应的状态码为200 OK
		// 设置Content-Type为image/png，表示响应体是一个PNG格式的图片
		// 并将生成的二维码图片数据作为响应体发送给客户端
		c.Data(http.StatusOK, "image/png", png)
	} else {
		// 如果"content"查询参数为空，则返回400 Bad Request状态码
		c.Status(http.StatusBadRequest)
	}
}
