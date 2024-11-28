package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ResponseType struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		City string `json:"city"`
		Data []struct {
			AirQuality  string `json:"air_quality"`
			Date        string `json:"date"`
			Temperature string `json:"temperature"`
			Weather     string `json:"weather"`
			Wind        string `json:"wind"`
		} `json:"data"`
	} `json:"data"`
}

func get_current_weather(args string) string {
	var result map[string]string
	// 解析 JSON 字符串
	err := json.Unmarshal([]byte(args), &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)

	}
	// 获取 city_name 的值
	cityName, exists := result["city_name"]
	if !exists {
		fmt.Println("city_name not found")
	}
	url := fmt.Sprintf("https://v2.api-m.com/api/weather?city=%s", cityName)
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
	var respdata ResponseType

	if err := json.Unmarshal(body, &respdata); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	// 用来存储所有天气信息的切片
	var weatherInfo []string
	// 遍历 data 数组并将每个对象的内容格式化为字符串
	for _, item := range respdata.Data.Data {
		result := fmt.Sprintf("Date: %s, Weather: %s, Temperature: %s, Wind: %s, Air Quality: %s",
			item.Date, item.Weather, item.Temperature, item.Wind, item.AirQuality)
		weatherInfo = append(weatherInfo, result)
	}
	// 将切片转换为单一字符串
	return strings.Join(weatherInfo, "\n")
}

func main() {
	gpt_model := "gpt-4-0613"
	OPENAI_API_KEY := "sk-xdvfSI0uUT7Zc8qiE82319E00cBb46C0917c603707A37378"
	OPENAI_BASE_URL := "https://one.api4gpt.com/v1"
	config := openai.DefaultConfig(OPENAI_API_KEY)
	config.BaseURL = OPENAI_BASE_URL
	client := openai.NewClientWithConfig(config)

	Tool := []openai.Tool{
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "get_current_weather",
				Description: "获取当前天气状况",
				Parameters: jsonschema.Definition{
					Type: jsonschema.Object,
					Properties: map[string]jsonschema.Definition{
						"city_name": {
							Type:        jsonschema.String,
							Description: "中国城市名称",
						},
					},
					Required: []string{"city_name"},
				},
			},
		},
	}
	Message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "不要对要插入函数的值做出假设。如果用户请求不明确，请要求澄清，必须用简体中文回答，请根据当前城市以及未来的天气状况",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "深圳今天天气怎么样?",
		},
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:      gpt_model,
			Messages:   Message,
			Tools:      Tool,
			ToolChoice: "auto",
		},
	)
	if err != nil {
		fmt.Printf("Error in ChatCompletion: %v\n", err)
		return
	}
	var messages []openai.ChatCompletionMessage
	if resp.Choices[0].FinishReason == "tool_calls" {
		//获取函数名称
		messages = append(messages, resp.Choices[0].Message)
		args := resp.Choices[0].Message.ToolCalls[0].Function.Arguments

		date := get_current_weather(args)
		ReplyMessage := openai.ChatCompletionMessage{
			Role:       openai.ChatMessageRoleTool,
			Content:    date,
			ToolCallID: resp.Choices[0].Message.ToolCalls[0].ID,
			Name:       resp.Choices[0].Message.ToolCalls[0].Function.Name,
		}
		roleMessage := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: "你是一位专业的天气预报员，首先告诉用户这几天的天气情况，然后再根据获取到的当前城市以及未来的天气状况，提醒用户需要注意什么，可以做什么",
		}
		messages = append(messages, ReplyMessage)
		messages = append(messages, roleMessage)

		resp, err = client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    gpt_model,
				Messages: messages,
			},
		)
		if err != nil {
			fmt.Printf("Error in ChatCompletion: %v\n", err)
			return
		}
		fmt.Println(resp.Choices[0].Message.Content)
	}
}
