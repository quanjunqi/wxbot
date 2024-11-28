package yuying

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

type Response struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Audiopath string `json:"audiopath"`
	APISource string `json:"api_source"`
}

func init() {
	engine := control.Register("fuli", &control.Options{
		Alias:    "fuli",
		Help:     "福利插件",
		Priority: 10,
	})
	fulimatch := []string{"骂我", "安慰我"}
	engine.OnFullMatchGroup(fulimatch).SetBlock(true).Handle(func(ctx *bot.Ctx) {

		// 判断匹配的内容
		switch ctx.Event.Message.Content {
		case "安慰我":
			data := yujie()
			ctx.ReplyImage(data)
		case "骂我":
			data := duiren()
			ctx.ReplyImage(data)
		default:
			text := "不行呀，现在有点忙"
			ctx.ReplyText(text)
		}
	})
}

// 御姐安慰
func yujie() string {
	url := "https://api.pearktrue.cn/api/yujie/?type=mp3"

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
	return response.Audiopath
}

// 怼人
func duiren() string {
	url := "https://api.pearktrue.cn/api/duiren/?type=mp3"

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
	return response.Audiopath
}
