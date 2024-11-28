package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	fmt.Println(string(responsebody))
	return respdata.Imgurl
}

func main() {
	age := 10
	file := "https://wechat-qjq.oss-cn-shenzhen.aliyuncs.com/6b6cd7a2f9f3afea4ee671c6f67db337.jpg"
	faceold(age, file)
}
