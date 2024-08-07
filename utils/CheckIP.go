package utils

import (
	"OneDrive-Download-Proxy/types"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

func CheckIP(ip string) (string, string) {
	urlConfig, _ := LoadConfig("./config/config.json")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ip.sb/geoip/"+ip, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("%s\n", bodyText)

	var ipInfoType types.IPInfoType
	err = json.Unmarshal(bodyText, &ipInfoType)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(ipInfoType.Isp)
	slog.Info("CheckIP", "ISP", ipInfoType.Isp)
	netIsp := ipInfoType.Isp

	if strings.Contains(netIsp, "China") && strings.Contains(netIsp, "Mobile") {
		return ipInfoType.Isp, urlConfig.URLConfig.CnMobile.URL
	} else if strings.Contains(netIsp, "China") && strings.Contains(netIsp, "Unicom") {
		return ipInfoType.Isp, urlConfig.URLConfig.CnUni.URL
	} else if strings.Contains(netIsp, "China") && strings.Contains(netIsp, "Telecom") {
		return ipInfoType.Isp, urlConfig.URLConfig.CnTele.URL
	}
	return ipInfoType.Isp, urlConfig.URLConfig.Other.URL
}
