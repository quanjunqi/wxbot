package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 定义响应结构体
type Response struct {
	Success bool   `json:"success"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	URL     string `json:"url"`
}

func main() {
	url := "https://api.vvhan.com/api/moyu?type=json"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		log.Println("请求错误:", err)

	}
	defer resp.Body.Close() // 确保在函数退出时关闭响应体

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Printf("请求失败，状态码：%d\n", resp.StatusCode)

	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应体错误:", err)

	}

	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("解析 JSON 错误:", err)

	}

	// 返回 URL 字段
	fmt.Println(response.URL)

}
