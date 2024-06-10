package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.cwl.gov.cn/")
	time.Sleep(2 * time.Second) // 等待页面加载完成

	cookies := page.MustCookies()
	fmt.Printf("Cookie %s\n", cookies)
	var cookieStr strings.Builder
	for _, cookie := range cookies {
		cookieStr.WriteString(cookie.Name)
		cookieStr.WriteString("=")
		cookieStr.WriteString(cookie.Value)
		cookieStr.WriteString(";")
	}

	// 打印拼接的 Cookie 字符串
	finalCookieStr := cookieStr.String()
	if len(finalCookieStr) > 0 {
		finalCookieStr = finalCookieStr[:len(finalCookieStr)-2] // 去掉最后的 "; "
	}
	// 解析后的格式存在问题，无法正常使用
	fmt.Printf("Cookie String: %s\n", finalCookieStr)
}
