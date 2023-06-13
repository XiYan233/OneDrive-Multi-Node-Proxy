package main

import (
	"OneDrive-Download-Proxy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
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

		if remoteIP == "" {
			ip = forwardedIp
		} else {
			ip = remoteIP
		}

		var dbPath = "./xdb/ip2region.xdb"
		searcher, err := xdb.NewWithFileOnly(dbPath)
		if err != nil {
			fmt.Printf("failed to create searcher: %s\n", err.Error())
			return
		}

		defer searcher.Close()

		// 移动  223.104.76.193
		// 联通  112.64.0.2
		// 电信  58.87.64.5
		// 广电  117.120.128.0

		fmt.Println(remoteIP)
		fmt.Println(ip)
		if err != nil {
			fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
			return
		}
		url := c.Request.URL.String()
		netStr, redirectUrl := utils.CheckIP(ip)
		fmt.Printf("{url:%s}\n", url)
		fmt.Printf("请求运营商为：%s，返回地址为：%s\n", netStr, redirectUrl+url)

		c.Redirect(http.StatusFound, redirectUrl+url)
	})

	r.Run(":8080")

}
