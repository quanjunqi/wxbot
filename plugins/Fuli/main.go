package Fuli

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

type Response struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    string `json:"data"`
	Video   string `json:"video"`
	Raw_url string `json:"raw_url"`
}

func init() {
	engine := control.Register("fuli", &control.Options{
		Alias:    "fuli",
		Help:     "福利插件",
		Priority: 10,
	})
	fulimatch := []string{"看美女", "福利视频"}
	engine.OnFullMatchGroup(fulimatch).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		date := randomFuli()
		ctx.ReplyImage(date)
		if date == "" {
			text := "看美女失败"
			ctx.ReplyText(text)
		}
	})
}

// 随机选择接口
func randomFuli() string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 将函数存储在切片中
	functions := []func() (string, bool){
		fuli1,
		fuli2,
		fuli3,
	}
	// 尝试执行函数，直到成功
	for len(functions) > 0 {
		randomIndex := rand.Intn(len(functions))
		result, success := functions[randomIndex]()
		if success {
			return result // 成功执行，返回结果
		}
	}
	return "All functions failed." // 所有函数都失败的情况
}

// 美女视频1
func fuli1() (string, bool) {
	url := "https://api.kuleu.com/api/xjj?type=json"
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
	if response.Code != 200 {
		return "", false
	}
	return response.Video, true
}

// 美女视频2
func fuli2() (string, bool) {
	url := "https://onexiaolaji.cn/RandomPicture/api/api-video.php"
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
	if response.Raw_url == "" {
		return "", false
	}
	return response.Raw_url, true
}

// 美女视频3
func fuli3() (string, bool) {
	url := "https://v2.api-m.com/api/meinv"
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
	if response.Code != 200 {
		return "", false
	}
	return response.Data, true
}
