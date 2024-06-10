package main

import (
	"fmt"
	"github.com/go-rod/rod"
)

func main() {
	// 启动浏览器
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// 打开一个新页面
	page := browser.MustPage("https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=2&issueStart=&issueEnd=&dayStart=&dayEnd=&pageNo=1&pageSize=30&week=&systemType=PC")

	// 等待页面加载完成
	page.MustWaitLoad()

	// 获取页面内容
	body, err := page.HTML()
	if err != nil {
		fmt.Println("Error getting page HTML:", err)
		return
	}

	// 打印页面内容
	fmt.Println(body)
}
