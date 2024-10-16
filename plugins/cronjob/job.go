package cronjob

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// 定义响应结构体
type Response struct {
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Success bool   `json:"success"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	URL     string `json:"url"`
}

func Moyu() string {
	url := "https://api.vvhan.com/api/moyu?type=json"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		log.Println("请求错误:", err)
		return "" // 返回空字符串
	}
	defer resp.Body.Close() // 确保在函数退出时关闭响应体

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Printf("请求失败，状态码：%d\n", resp.StatusCode)
		return "" // 返回空字符串
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应体错误:", err)
		return "" // 返回空字符串
	}

	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("解析 JSON 错误:", err)
		return "" // 返回空字符串
	}

	// 返回 URL 字段
	return response.URL
}

// 舔狗日记
func Tiangou() string {
	url := "https://api.dzzui.com/api/tiangou?format=json"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		log.Println("请求错误:", err)
		return "" // 返回空字符串
	}
	defer resp.Body.Close() // 确保在函数退出时关闭响应体

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Printf("请求失败，状态码：%d\n", resp.StatusCode)
		return "" // 返回空字符串
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应体错误:", err)
		return "" // 返回空字符串
	}

	// 解析 JSON 响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("解析 JSON 错误:", err)
		return "" // 返回空字符串
	}

	// 返回 URL 字段
	return response.Text
}
