package chatroom

import (
	"fmt"
	"regexp"
	"strings"
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

func init() {
	engine := control.Register("chatroom", &control.Options{
		Alias:    "chatroom",
		Help:     "群聊插件",
		Priority: 20,
	})
	engine.OnMessage(bot.OnlyGroup).SetBlock(false).Handle(func(ctx *bot.Ctx) {
		//存储用户拍一拍状态
		PatMap := make(map[string]int)

		if ctx.Ispat() {
			re := regexp.MustCompile(`(.*?)拍了拍\s*(.*)`)
			if matches := re.FindStringSubmatch(ctx.Event.Message.Content); len(matches) > 2 {
				// 检查该用户是否已经拍过
				// 生成唯一的键
				after := matches[2]
				if strings.Contains(after, "我") {

					key := ctx.Event.FromGroup + ctx.Event.FromWxId
					if PatMap[key] == 0 { // 第一次拍了拍
						ctx.ReplyPat(ctx.Event.FromWxId)
						PatMap[key]++ // 更新该用户的状态
					} else {
						ctx.ReplyPat(ctx.Event.FromWxId)
						ctx.ReplyText("别拍我没结果") // 不是第一次
					}
				} else {
					fmt.Println("没有找到'拍了拍'")
				}
			}
		}
	})
}
