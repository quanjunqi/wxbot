package alipay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

type Response struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Number    string `json:"number"`
	Audiourl  string `json:"audiourl"`
	APISource string `json:"api_source"`
}

func init() {
	engine := control.Register("alipay", &control.Options{
		Alias:    "alipay",
		Help:     "支付宝到账语音生成",
		Priority: 10,
	})
	engine.OnRegex(`支付宝到账语音生成\s*(.*)`).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		// 获取匹配的内容
		matched := ctx.State["regex_matched"].([]string)
		if len(matched) > 1 {
			matchedNumber := matched[1] // 获取捕获的数字字符串

			// 将字符串转换为整数
			number, err := strconv.Atoi(matchedNumber)
			if err != nil {
				text := "数字转换失败"
				ctx.ReplyText(text)
				return
			}

			date := alipay(number) // 使用整数调用 alipay 函数
			ctx.ReplyImage(date)
			if date == "" {
				text := "生成失败"
				ctx.ReplyText(text)
			}
		} else {
			text := "没有匹配到数字"
			ctx.ReplyText(text)
		}
	})
}
func alipay(number int) string {
	url := fmt.Sprintf("https://api.pearktrue.cn/api/alipay/?number=%d&type=json", number)
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
	if response.Code == 200 {
		return response.Audiourl
	}
	// 返回 URL 字段
	return ""

}
