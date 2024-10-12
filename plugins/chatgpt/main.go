package chatgpt

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
	"wxbot/engine/bot"
	"wxbot/engine/control"
	"wxbot/engine/pkg/sqlite"
)

// //go:embeddata
// var chatGptData embed.FS

var (
	db          sqlite.DB // 数据库
	chatRoomCtx sync.Map  // 聊天室消息上下文
)

// ChatRoom chatRoomCtx -> ChatRoom => 维系每个人的上下文
type ChatRoom struct {
	chatId   string    // 聊天室ID, 格式为: 聊天室ID_发送人ID
	chatTime time.Time // 聊天时间
	content  []Message // 聊天上下文内容
}

func init() {

	engine := control.Register("chatgpt", &control.Options{
		Alias:    "chatgpt",
		Help:     "智能对话",
		Priority: 0,
	})
	if err := sqlite.Open("/root/wxbot/plugins/chatgpt/chatgpt.db", &db); err != nil {
		log.Fatalf("open sqlite db failed: %v", err)
	}

	engine.OnMessage(bot.OnlyAtMe).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		var (
			now = time.Now().Local()
			msg = ctx.MessageString()

			chatRoom = ChatRoom{
				chatId:   fmt.Sprintf("%s_%s", ctx.Event.FromUniqueID, ctx.Event.FromWxId),
				chatTime: time.Now().Local(),
				content:  []Message{},
			}
		)
		// 正式处理
		if c, ok := chatRoomCtx.Load(chatRoom.chatId); ok {
			// 判断距离上次聊天是否超过10分钟了
			if now.Sub(c.(ChatRoom).chatTime) > 10*time.Minute {
				chatRoomCtx.LoadAndDelete(chatRoom.chatId)
				chatRoom.content = []Message{{Role: "user", Content: msg}}
			} else {
				chatRoom.content = append(c.(ChatRoom).content, Message{Role: "user", Content: msg})
			}
		} else {
			chatRoom.content = []Message{{Role: "user", Content: msg}}
		}
		replyMessage := Chatgpt_text(chatRoom.content)
		chatRoom.content = append(chatRoom.content, Message{Role: "assistant", Content: replyMessage.ReplyContent})
		chatRoomCtx.Store(chatRoom.chatId, chatRoom)
		if replyMessage.Replytext != "" {
			// 根据换行符分割文本
			lines := strings.Split(replyMessage.Replytext, "\n")

			// // 删除最后一行商品链接
			// if len(lines) > 0 {
			// 	lines = lines[:len(lines)-1]
			// }
			// 分段发送结果
			for i, line := range lines {
				// 模拟随机延迟
				time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)

				// 第一行使用 ReplyTextAt，后续行使用 ReplyText
				if i == 0 {
					ctx.ReplyTextAt(line)
				} else {
					ctx.ReplyText(line)
				}
			}

		} else {
			ctx.ReplyTextAt(replyMessage.ReplyContent)
		}
		if len(replyMessage.Replyurl) > 0 {
			for _, url := range replyMessage.Replyurl {
				time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
				ctx.ReplyImage(url)
			}
		}
	})
}
