package main

import (
	"OneDrive-Download-Proxy/utils"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/*path", func(c *gin.Context) {
		ip := c.ClientIP()
		var dbPath = "./xdb/ip2region.xdb"
		searcher, err := xdb.NewWithFileOnly(dbPath)
		if err != nil {
			fmt.Printf("failed to create searcher: %s\n", err.Error())
			return
		}

		defer searcher.Close()

		// do the search
		// var ip = "1.2.3.4"
		// 移动  223.104.76.193
		// 联通  112.64.0.2
		// 电信  58.87.64.5
		//ip = "58.87.64.5"
		var tStart = time.Now()
		region, err := searcher.SearchByStr(ip)
		if err != nil {
			fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
			return
		}
		url := c.Request.URL.String()
		redirectUrl := utils.CheckIP(region) + url
		fmt.Printf("{url:%s}\n", url)
		fmt.Printf("{ip:%s,region: %s, took: %s}\n", ip, region, time.Since(tStart))
		fmt.Printf("请求运营商为：%s，返回地址为：%s\n", region, redirectUrl)

		c.Redirect(http.StatusFound, redirectUrl)
	})

	// 为HTTPS配置TLS证书和密钥
	server := &http.Server{
		Addr:    ":443",
		Handler: r,
		TLSConfig: &tls.Config{
			// 配置TLS证书和密钥的路径
			Certificates: []tls.Certificate{loadTLSCert("./SSL/cert.pem", "./SSL/key.pem")},
		},
	}

	// 启动HTTPS服务器
	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
}

func loadTLSCert(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	return cert
}
