package utils

import "strings"

func CheckIP(str string) string {
	urlConfig, _ := LoadConfig("./config/config.json")
	if strings.Contains(str, "中国") && strings.Contains(str, "移动") {
		return urlConfig.URLConfig.CnMobile.URL
	} else if strings.Contains(str, "中国") && strings.Contains(str, "联通") {
		return urlConfig.URLConfig.CnUni.URL
	} else if strings.Contains(str, "中国") && strings.Contains(str, "电信") {
		return urlConfig.URLConfig.CnTele.URL
	}
	return urlConfig.URLConfig.Other.URL
}
