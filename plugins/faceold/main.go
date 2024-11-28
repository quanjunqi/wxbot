package faceold

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

type Response struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Imgurl    string `json:"imgurl"`
	APISource string `json:"api_source"`
}
type RequestType struct {
	File string `json:"file"`
	Age  int    `json:"age"`
}

func init() {
	engine := control.Register("fuli", &control.Options{
		Alias:    "fuli",
		Help:     "年龄转换插件",
		Priority: 10,
	})
	engine.OnRegex(`年龄转换\s*(\d+)`).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		if ctx.Event.ReferenceMessage != nil {
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
				date := faceold(number, ctx.Event.ReferenceMessage.ImageURL) // 使用整数调用
				ctx.ReplyImage(date)
				if date == "" {
					text := "生成失败"
					ctx.ReplyText(text)
				}
			} else {
				text := "没有匹配到数字"
				ctx.ReplyText(text)
			}
		}
	})
}
func faceold(age int, file string) string {
	url := "https://api.pearktrue.cn/api/faceold/"
	request := RequestType{
		File: file,
		Age:  age,
	}
	requestdata, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(requestdata))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responsebody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var respdata Response

	if err := json.Unmarshal(responsebody, &respdata); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	return respdata.Imgurl
}
