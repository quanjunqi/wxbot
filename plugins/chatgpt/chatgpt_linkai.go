package chatgpt

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Chatgpt_text(messages []Message) ReplyMessage {
	client := &http.Client{}
	// 创建消息
	// messages := []Message{
	// 	{
	// 		Role:    "user",
	// 		Content: text,
	// 	},
	// }
	var chatMessages []Message

	chatMessages = append(chatMessages, messages...)
	data := RequestData{
		AppCode:  "QdE4kjamIg",
		Messages: chatMessages,
	}
	requestdata, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "https://api.link-ai.tech/v1/chat/completions", bytes.NewBuffer(requestdata))
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
	reply := ReplyMessage{
		ReplyContent: respdata.Choices[0].Message.Content,
		Replyurl:     respdata.Choices[0].ImgUrls,
		Replytext:    respdata.Choices[0].TextContent,
	}

	return reply
}
