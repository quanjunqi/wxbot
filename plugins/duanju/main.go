package duanju

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

func init() {
	engine := control.Register("duanju", &control.Options{
		Alias:    "duanju",
		Help:     "短剧插件",
		Priority: 10,
	})
	engine.OnRegex(`搜短剧\s*(.*)`).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		duanjuName := ctx.State["regex_matched"].([]string)[1]
		apiItems := Duanju(duanjuName)
		if apiItems == nil {
			text := "未找到对应的短剧信息，缩减关键词试试"
			ctx.ReplyText(text)
			return
		}
		// 用于存储所有的消息内容
		var allMessages string
		maxItems := 15 // 设置最大返回条数

		for i, item := range apiItems {
			if i >= maxItems {
				break // 如果已收集到最大条数，则退出循环
			}
			text := fmt.Sprintf("短剧:%s\n下载地址:%s\n ", item.Title, item.URL)
			allMessages += text // 拼接消息
		}
		ctx.ReplyText(allMessages)
	})
}

// 定义短剧结构体以匹配 JSON 数据的结构
type apiResponse struct {
	Auth    string    `json:"auth"`
	Success bool      `json:"success"`
	Total   int       `json:"total"`
	Data    []apiItem `json:"data"`
}

type apiItem struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Episodes int    `json:"episodes"`
}

const key = 333

// 短剧搜索
func Duanju(search string) []apiItem {
	client := &http.Client{}
	url := fmt.Sprintf("https://api-duanju.cooks.team/?key=%s&pw=%d", search, key) // 修改为 %s 以匹配密钥类型
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 解析 JSON 数据
	var response apiResponse
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	if response.Success && response.Total > 0 {
		return response.Data
	}
	return nil
}
