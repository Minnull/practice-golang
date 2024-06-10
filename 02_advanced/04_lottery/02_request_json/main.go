package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/go-rod/rod"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 定义响应数据结构
type Response struct {
	State    int      `json:"state"`
	Message  string   `json:"message"`
	Total    int      `json:"total"`
	PageNum  int      `json:"pageNum"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
	Tflag    int      `json:"Tflag"`
	Result   []Result `json:"result"`
}

// 定义Result结构
type Result struct {
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	DetailsLink string       `json:"detailsLink"`
	VideoLink   string       `json:"videoLink"`
	Date        string       `json:"date"`
	Week        string       `json:"week"`
	Red         string       `json:"red"`
	Blue        string       `json:"blue"`
	Blue2       string       `json:"blue2"`
	Sales       string       `json:"sales"`
	Poolmoney   string       `json:"poolmoney"`
	Content     string       `json:"content"`
	Addmoney    string       `json:"addmoney"`
	Addmoney2   string       `json:"addmoney2"`
	Msg         string       `json:"msg"`
	Z2add       string       `json:"z2add"`
	M2add       string       `json:"m2add"`
	Prizegrades []Prizegrade `json:"prizegrades"`
}

// 定义Prizegrade结构
type Prizegrade struct {
	Type      int    `json:"type"`
	Typenum   string `json:"typenum"`
	Typemoney string `json:"typemoney"`
}

func main() {
	url := "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=2&issueStart=&issueEnd=&dayStart=&dayEnd=&pageNo=1&pageSize=30&week=&systemType=PC"

	// 创建一个请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 添加请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("sec-ch-ua", `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Referer", "https://www.cwl.gov.cn/ygkj/kjgg/")
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")

	// 设置Cookie，cookie每次得对上，后面更新cookie
	// 需要从页面找到正确的cookie，然后设置到这里
	req.Header.Set("Cookie", "HMF_CI=863b8f711aadb49c5ad8e3c4d359d2ade8f9adfaa99602664aa7e767b65fe5659f85cd6d0e4c507019e3482d99ac230e586dd26ee10fab57de4b63bb41f193e050; 21_vq=2")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// 读取和解压响应体
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// 读取解压后的响应体
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印解压后的响应体
	fmt.Println("Response Body:", string(body))

	// 解析JSON数据
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印解析后的数据
	fmt.Printf("State: %d\n", response.State)
	fmt.Printf("Message: %s\n", response.Message)
	fmt.Printf("Total: %d\n", response.Total)
	fmt.Printf("PageNum: %d\n", response.PageNum)
	fmt.Printf("PageNo: %d\n", response.PageNo)
	fmt.Printf("PageSize: %d\n", response.PageSize)
	fmt.Printf("Tflag: %d\n", response.Tflag)

	for _, result := range response.Result {
		fmt.Printf("Name: %s, Code: %s, Date: %s, Red: %s, Blue: %s, Sales: %s, Poolmoney: %s\n", result.Name, result.Code, result.Date, result.Red, result.Blue, result.Sales, result.Poolmoney)
		for _, prize := range result.Prizegrades {
			fmt.Printf("Type: %d, Typenum: %s, Typemoney: %s\n", prize.Type, prize.Typenum, prize.Typemoney)
		}
	}
}

func getCookie() string {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.cwl.gov.cn/")
	time.Sleep(2 * time.Second) // 等待页面加载完成

	cookies := page.MustCookies()
	var cookieStr strings.Builder
	for _, cookie := range cookies {
		cookieStr.WriteString(cookie.Name)
		cookieStr.WriteString("=")
		cookieStr.WriteString(cookie.Value)
		cookieStr.WriteString("; ")
	}

	// 打印拼接的 Cookie 字符串
	finalCookieStr := cookieStr.String()
	if len(finalCookieStr) > 0 {
		finalCookieStr = finalCookieStr[:len(finalCookieStr)-2] // 去掉最后的 "; "
	}
	return finalCookieStr
}
