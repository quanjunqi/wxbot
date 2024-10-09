package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

func init() {
	engine := control.Register("baidubaike", &control.Options{
		Alias:    "chatgpt",
		Help:     "测试",
		Priority: 0,
	})

	engine.OnMessage(bot.OnlyAtMe).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		msg := ctx.MessageString()
		text := chatgpt_text(msg)
		ctx.ReplyTextAt(text)
	})
}

func chatgpt_text(text string) string {
	client := &http.Client{}
	// 创建消息
	messages := []Message{
		{
			Role:    "user",
			Content: text,
		},
	}

	data := RequestData{
		AppCode:  "FlF1XQ9N",
		Messages: messages,
	}
	requestdata, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "https://api.link-ai.chat/v1/chat/completions", bytes.NewBuffer(requestdata))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer Link_lErYDs0p4AntBjjUHhrZTbM7bHom54Q0SseKi1jrjT")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var respdata ResponseData

	if err := json.Unmarshal(bodyText, &respdata); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}
	contont := respdata.Choices[0].Message.Content
	fmt.Println(contont)
	return contont
}
