package test

import (
	"wxbot/engine/bot"
	"wxbot/engine/control"
)

func init() {
	engine := control.Register("baidubaike", &control.Options{
		Alias: "test",
		Help:  "测试",
	})

	engine.OnMessage(bot.OnlyAtMe).SetBlock(true).Handle(func(ctx *bot.Ctx) {
		if ctx.Event.IsAtMe {
			ctx.ReplyText("你好")
		}
	})
}
