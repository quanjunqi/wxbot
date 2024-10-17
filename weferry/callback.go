package weferry

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"wxbot/engine/bot"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const (
	eventTextMessage      = 1     // 文字消息
	eventImageMessage     = 3     // 图片消息
	eventSystemMessage    = 10000 //x系统消息
	eventReferenceMessage = 49    //引用消息

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
	fmt.Println(string(recv))
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
					Id:      gjson.Get(resp, "id").Int(),
					Type:    bot.MsgTypeText,
					Content: gjson.Get(resp, "content").String(),
				},
				IsAtMe: strings.Contains(gjson.Get(resp, "content").String(), "@小帅"),
			}
		case eventImageMessage: //图片消息
			event = bot.Event{
				Type:           bot.EventGroupChat,
				FromUniqueID:   gjson.Get(resp, "roomid").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromGroup:      gjson.Get(resp, "roomid").String(),
				FromGroupName:  GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				Message: &bot.Message{
					Id:      gjson.Get(resp, "id").Int(),
					Type:    bot.MsgTypeImage,
					Content: gjson.Get(resp, "content").String(),
				},
			}
			path := "C:/Users/Administrator/Desktop/image"
			time.Sleep(3 * time.Second)
			SaveImage(event.Message.Id, path, gjson.Get(resp, "extra").String())
		case eventSystemMessage: //系统消息消息
			event = bot.Event{
				Type:           bot.EventGroupChat,
				FromUniqueID:   gjson.Get(resp, "roomid").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromGroup:      gjson.Get(resp, "roomid").String(),
				FromGroupName:  GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				Message: &bot.Message{
					Id:      gjson.Get(resp, "id").Int(),
					Type:    bot.MsgTypeSystem,
					Content: gjson.Get(resp, "content").String(),
				},
			}
		case eventReferenceMessage: //引用消息
			event = bot.Event{
				Type:           bot.EventGroupChat,
				FromUniqueID:   gjson.Get(resp, "roomid").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromGroup:      gjson.Get(resp, "roomid").String(),
				FromGroupName:  GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				IsAtMe:         strings.Contains(gjson.Get(resp, "content").String(), "@小帅"),
				Message: &bot.Message{
					Id: gjson.Get(resp, "id").Int(),
				},
			}
			var refer ReferenceXml
			if err := xml.Unmarshal([]byte(gjson.Get(resp, "content").String()), &refer); err == nil {
				if refer.Appmsg.Refermsg != nil { // 引用消息
					event.Message.Type = bot.MsgTypeText // 方便匹配
					event.Message.Content = refer.Appmsg.Title
					event.ReferenceMessage = &bot.ReferenceMessage{
						ReferenceMessageType: bot.MsgTypeText,
						ReferenceMessageID:   refer.Appmsg.Refermsg.Svrid,
						FromUser:             refer.Appmsg.Refermsg.Fromusr,
						ChatUser:             refer.Appmsg.Refermsg.Chatusr,
						Content:              refer.Appmsg.Refermsg.Content,
					}
				}
				if refer.Appmsg.Refermsg.Type == 3 {
					event.ReferenceMessage.ReferenceMessageType = bot.MsgTypeImage
				}
			}

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
					Id:      gjson.Get(resp, "id").Int(),
					Type:    bot.MsgTypeText,
					Content: gjson.Get(resp, "content").String(),
				},
				// IsAtMe: true,
			}
		case eventImageMessage: //图片消息
			event = bot.Event{
				Type:           bot.EventPrivateChat,
				FromUniqueID:   gjson.Get(resp, "sender").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "sender").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				Message: &bot.Message{
					Id:      gjson.Get(resp, "id").Int(),
					Type:    bot.MsgTypeImage,
					Content: gjson.Get(resp, "content").String(),
				},
			}
			path := "C:/Users/Administrator/Desktop/image"
			time.Sleep(3 * time.Second)
			SaveImage(event.Message.Id, path, gjson.Get(resp, "extra").String())

		case eventReferenceMessage: //引用消息
			event = bot.Event{
				Type:           bot.EventPrivateChat,
				FromUniqueID:   gjson.Get(resp, "roomid").String(),
				FromUniqueName: GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromGroup:      gjson.Get(resp, "roomid").String(),
				FromGroupName:  GetChatRoomNick(gjson.Get(resp, "roomid").String()),
				FromWxId:       gjson.Get(resp, "sender").String(),
				FromName:       GetChatRoomNick(gjson.Get(resp, "sender").String()),
				Message: &bot.Message{
					Id: gjson.Get(resp, "id").Int(),
				},
			}
			var refer ReferenceXml
			if err := xml.Unmarshal([]byte(gjson.Get(resp, "content").String()), &refer); err == nil {
				if refer.Appmsg.Refermsg != nil { // 引用消息
					event.Message.Type = bot.MsgTypeText // 方便匹配
					event.Message.Content = refer.Appmsg.Title
					event.ReferenceMessage = &bot.ReferenceMessage{
						ReferenceMessageType: bot.MsgTypeText,
						ReferenceMessageID:   refer.Appmsg.Refermsg.Svrid,
						FromUser:             refer.Appmsg.Refermsg.Fromusr,
						ChatUser:             refer.Appmsg.Refermsg.Chatusr,
						Content:              refer.Appmsg.Refermsg.Content,
					}
				}
				if refer.Appmsg.Refermsg.Type == 3 {
					event.ReferenceMessage.ReferenceMessageType = bot.MsgTypeImage
				}
			}
		}
	}
	return &event
}
