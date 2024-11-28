package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestData struct {
	AppCode  string    `json:"app_code"`
	Messages []Message `json:"messages"`
}

type ResponseData struct {
	Choices []Choices   `json:"choices"`
	Usage   Usage       `json:"usage"`
	Model   interface{} `json:"model"`
	TraceID string      `json:"trace_id"`
	Agent   Agent       `json:"agent"`
}
type Message struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal"`
}
type FinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}
type Choices struct {
	Index         int           `json:"index"`
	Message       Message       `json:"message"`
	Logprobs      interface{}   `json:"logprobs"`
	FinishDetails FinishDetails `json:"finish_details"`
}
type PromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}
type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}
type Usage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}
type Chain struct {
	PluginName  string `json:"plugin_name"`
	PluginIcon  string `json:"plugin_icon"`
	PluginInput string `json:"plugin_input"`
}
type Agent struct {
	Status          string  `json:"status"`
	Chain           []Chain `json:"chain"`
	NeedShowPlugin  bool    `json:"need_show_plugin"`
	NeedShowThought bool    `json:"need_show_thought"`
}

func main1() {
	client := &http.Client{}
	// 创建消息
	messages := []Message{
		{
			Role:    "user",
			Content: "搜索沁园春雪的整首词完整版",
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
	fmt.Println(string(bodyText))
	contont := respdata.Choices[0].Message.Content
	fmt.Println(contont)
}
