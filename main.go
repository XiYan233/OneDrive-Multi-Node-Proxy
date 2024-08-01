package main

import (
	"OneDrive-Download-Proxy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/*path", func(c *gin.Context) {
		remoteAddr := c.Request.RemoteAddr
		remoteIP, _, _ := net.SplitHostPort(remoteAddr)
		forwardedIp := c.Request.Header.Get("X-Forwarded-For")

		ip := ""

		if forwardedIp == "" {
			ip = remoteIP
		} else {
			ip = forwardedIp
		}

		// 移动  223.104.76.193
		// 联通  112.64.0.2
		// 电信  58.87.64.5
		// 广电  117.120.128.0

		slog.Info("Main", "RemoteIP", remoteIP)
		slog.Info("Main", "IP", ip)
		url := c.Request.URL.String()
		netStr, redirectUrl := utils.CheckIP(ip)
		//netStr, redirectUrl := utils.CheckIP("117.120.128.0")
		fmt.Printf("{url:%s}\n", url)
		fmt.Printf("请求运营商为：%s，返回地址为：%s\n", netStr, redirectUrl+url)

		c.Redirect(http.StatusFound, redirectUrl+url)
	})

	r.Run(":8080")

}
