package bot

import (
	"log"

	"github.com/gin-gonic/gin"
)

type IFramework interface { //定义机器人的接口
	//回调接口
	Callback(*gin.Context, func(*Event, IFramework))
	SendText(receiver, msg string) error
	SendTextAt(receiver, msg, aters string) error
	SendImage(receiver, path string) error
	SendPat(roomid, wxid string) error
	GetChatRoomNumber(roomid string) int
	GetChatRoomNick(userNameId string) string
}

// SendText 发送文本消息到指定好友
func (ctx *Ctx) SendText(receiver, msg string) error {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if msg == "" {
		return nil
	}
	return ctx.framework.SendText(receiver, msg)
}

// SendText 发送文本消息并At指定好友
func (ctx *Ctx) SendTextAt(receiver, msg, aters string) error {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if msg == "" && aters == "" {
		return nil
	}
	return ctx.framework.SendTextAt(receiver, msg, aters)
}

// SendImage 发送图片消息到指定好友
func (ctx *Ctx) SendImage(receiver, path string) error {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if path == "" {
		return nil
	}
	return ctx.framework.SendImage(receiver, path)
}

// SendPat 发送拍一拍消息到指定好友
func (ctx *Ctx) SendPat(roomid, wxid string) error {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if roomid == "" || wxid == "" {
		return nil
	}
	return ctx.framework.SendPat(roomid, wxid)
}

// GetChatRoomNumber获取群人数
func (ctx *Ctx) GetChatRoomNumber(roomid string) int {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if roomid == "" {
		log.Fatal("群ID不能为空")
	}
	return ctx.framework.GetChatRoomNumber(roomid)
}

// GetChatRoomNumber获取群成员昵称
func (ctx *Ctx) GetChatRoomNick(userNameId string) string {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if userNameId == "" {
		log.Fatal("群成员ID不能为空")
	}
	return ctx.framework.GetChatRoomNick(userNameId)
}

// ReplyImage  回复图片消息
func (ctx *Ctx) ReplyImage(path string) error {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if path == "" {
		return nil
	}
	return ctx.framework.SendImage(ctx.Event.FromUniqueID, path)
}

// ReplyText 回复文本消息
func (ctx *Ctx) ReplyText(text string) error {
	if text == "" {
		return nil
	}
	return ctx.SendText(ctx.Event.FromUniqueID, text)
}

// ReplyPat 回复拍一拍消息
func (ctx *Ctx) ReplyPat(wxid string) error {
	if wxid == "" {
		return nil
	}
	return ctx.SendPat(ctx.Event.FromUniqueID, wxid)
}

// ReplyText 回复文本消息并At指定好友
func (ctx *Ctx) ReplyTextAt(text string) error {
	if text == "" {
		return nil
	}
	return ctx.SendTextAt(ctx.Event.FromUniqueID, text, ctx.Event.FromWxId)
}
