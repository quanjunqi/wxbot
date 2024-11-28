package cronjob

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 定义响应结构体
type Response struct {
	Code    int    `json:"code"`
	Text    string `json:"text"`
	Success bool   `json:"success"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	URL     string `json:"url"`
	Msg     string `json:"msg"`
	Data    Data   `json:"data"`
}
type Data struct {
	GregorianDateTime string `json:"GregorianDateTime"`
	LunarDateTime     string `json:"LunarDateTime"`
	LunarShow         string `json:"LunarShow"`
	IsJieJia          string `json:"IsJieJia"`
	LJie              string `json:"LJie"`
	GJie              string `json:"GJie"`
	Yi                string `json:"Yi"`
	Ji                string `json:"Ji"`
	ShenWei           string `json:"ShenWei"`
	Taishen           string `json:"Taishen"`
	Chong             string `json:"Chong"`
	SuiSha            string `json:"SuiSha"`
	WuxingJiazi       string `json:"WuxingJiazi"`
	WuxingNaYear      string `json:"WuxingNaYear"`
	WuxingNaMonth     string `json:"WuxingNaMonth"`
	WuxingNaDay       string `json:"WuxingNaDay"`
	MoonName          string `json:"MoonName"`
	XingEast          string `json:"XingEast"`
	XingWest          string `json:"XingWest"`
	PengZu            string `json:"PengZu"`
	JianShen          string `json:"JianShen"`
	TianGanDiZhiYear  string `json:"TianGanDiZhiYear"`
	TianGanDiZhiMonth string `json:"TianGanDiZhiMonth"`
	TianGanDiZhiDay   string `json:"TianGanDiZhiDay"`
	LMonthName        string `json:"LMonthName"`
	LYear             string `json:"LYear"`
	LMonth            string `json:"LMonth"`
	LDay              string `json:"LDay"`
	SolarTermName     string `json:"SolarTermName"`
}

// 摸鱼混子
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

//万年历
func LaoHuangLi() string {
	now := time.Now()
	year, month, day := now.Date()
	url := fmt.Sprintf("https://www.36jxs.com/api/Commonweal/almanac?sun=%d-%d-%d", year, month, day)
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
	var text string
	if response.Msg == "成功获取" {
		text = fmt.Sprintf("公历：%s\n农历：%s\n宜：%s\n忌：%s\n神位：%s\n胎神：%s\n生肖吉冲：%s\n五行：%s\n星座：%s\n",
			response.Data.GregorianDateTime, response.Data.LunarDateTime, response.Data.Yi, response.Data.Ji, response.Data.ShenWei, response.Data.Taishen, response.Data.Chong, response.Data.WuxingNaDay, response.Data.XingWest)
	}
	// 返回 URL 字段
	return text

}
// kfc文案
func KFC() string {
	url := "https://api.pearktrue.cn/api/kfc/?type=json"

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
