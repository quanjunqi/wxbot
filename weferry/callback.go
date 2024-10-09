package weferry

import (
	"log"
	"net/http"
	"strings"

	"wxbot/engine/bot"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const (
	eventTextMessage  = 1 // 文字消息
	eventImageMessage = 3 // 图片
)

type Framework struct {
	BotWxId string // 机器人微信ID
	ApiUrl  string // http api地址
}

func New(botWxId, apiUrl string) *Framework {
	return &Framework{
		BotWxId: botWxId,
		ApiUrl:  apiUrl,
	}
}

func (f *Framework) Callback(ctx *gin.Context, handler func(*bot.Event, bot.IFramework)) {
	recv, err := ctx.GetRawData()
	if err != nil {
		log.Fatalf("[Dean] 接收回调错误, error: %v", err)
		return
	}
	// fmt.Println(string(recv))
	handler(buildEvent(string(recv)), f)
	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func buildEvent(resp string) *bot.Event {
	var event bot.Event
	switch gjson.Get(resp, "is_group").Bool() { //群聊
	case true: //群聊
		switch gjson.Get(resp, "type").Int() {
		case eventTextMessage: //文字消息
			event = bot.Event{
				Type:           bot.EventGroupChat,
				FromUniqueID:   gjson.Get(resp, "roomid").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromGroup:      gjson.Get(resp, "roomid").String(),
				FromGroupName:  GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				Message: &bot.Message{
					Id:      gjson.Get(resp, "id").String(),
					Type:    gjson.Get(resp, "type").Int(),
					Content: gjson.Get(resp, "content").String(),
				},
				IsAtMe: strings.Contains(gjson.Get(resp, "content").String(), "@小帅"),
			}
		case eventImageMessage: //图片消息
		}
	case false: //私聊
		switch gjson.Get(resp, "type").Int() {
		case eventTextMessage: //文字消息
			event = bot.Event{
				Type:           bot.EventPrivateChat,
				FromUniqueID:   gjson.Get(resp, "sender").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "sender").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				Message: &bot.Message{
					Id:      gjson.Get(resp, "id").String(),
					Type:    gjson.Get(resp, "type").Int(),
					Content: gjson.Get(resp, "content").String(),
				},
			}
		}
	}
	return &event
}
