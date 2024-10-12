package weferry

import (
	"encoding/json"
	"fmt"
)


// 发送文字消息
func (f *Framework) SendText(receiver, msg string) error {
	apiUrl := fmt.Sprintf("%s/text", f.ApiUrl)
	request := RequestType{
		Receiver: receiver,
		Msg:      msg,
	}
	requestdata, err := json.Marshal(request)
	if err != nil {
		return err
	}
	resp, err := sendHTTPRequest("POST", apiUrl, requestdata)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// 发送文字消息并At群成员
func (f *Framework) SendTextAt(receiver, msg, aters string) error {
	apiUrl := fmt.Sprintf("%s/text", f.ApiUrl)
	msg = fmt.Sprintf("@%s\t\t%s", GetChatRoomNick(aters), msg)
	request := RequestType{
		Receiver: receiver,
		Msg:      msg,
		Aters:    aters,
	}
	requestdata, err := json.Marshal(request)
	if err != nil {
		return err
	}
	resp, err := sendHTTPRequest("POST", apiUrl, requestdata)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// 发送图片消息
func (f *Framework) SendImage(receiver, path string) error {
	apiUrl := fmt.Sprintf("%s/image", f.ApiUrl)
	request := RequestType{
		Receiver: receiver,
		Path:     path,
	}
	requestdata, err := json.Marshal(request)
	if err != nil {
		return err
	}
	resp, err := sendHTTPRequest("POST", apiUrl, requestdata)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
