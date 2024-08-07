package controller

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取本地ip地址
func AddressesController(c *gin.Context) {
	addrList, _ := net.InterfaceAddrs()
	res := make([]string, 0)
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				res = append(res, ipNet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": res,
	})
}
