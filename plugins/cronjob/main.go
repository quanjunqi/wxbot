package cronjob

import (
	"sync"
	"time"
	"wxbot/cronjob"
	"wxbot/engine/bot"
	"wxbot/engine/control"
	"wxbot/weferry"
)

var (
	ApiUrl     = "http://124.220.212.132:10010"
	BootWxID   = ""
	once       sync.Once
	framework  bot.IFramework
	cronBuffer *cronjob.CronjobBuffer
)

func init() {
	// 确保这个初始化逻辑只执行一次
	control.Register("cronjob", &control.Options{
		Alias:    "cronjob",
		Help:     "定时任务",
		Priority: 10,
	})
	once.Do(func() {
		// 初始化 IFramework 实现
		framework = bot.IFramework(weferry.New(BootWxID, ApiUrl))

		// 创建 cronjob 缓冲区
		cronBuffer = cronjob.NewCronjobBuffer(4096)
		cronBuffer.Loop(5 * time.Second)

		// 添加定时任务
		cronBuffer.AddTask(cronjob.CronjoBufferItem{
			Function: func() {
				// framework.SendText("52386108522@chatroom", "下午好")
				framework.SendImage("52386108522@chatroom", Moyu())
			},
			CronExpression: "0 9 * * *", // 每天9:00执行,
			ID:             "摸鱼图片",
		})
		cronBuffer.AddTask(cronjob.CronjoBufferItem{
			Function: func() {
				framework.SendText("52386108522@chatroom", LaoHuangLi())
			},
			CronExpression: "0 10 * * *", // 每天10:00执行,
			ID:             "万年历",
		})		
		cronBuffer.AddTask(cronjob.CronjoBufferItem{
			Function: func() {
				framework.SendText("52386108522@chatroom", KFC())
			},
			CronExpression: "0 14 * * 4", // 每周四14:00执行,
			ID:             "kfc",
		})
		// cronBuffer.AddTask(cronjob.CronjoBufferItem{
		// 	Function: func() {
		// 		framework.SendText("52386108522@chatroom", Tiangou())
		// 	},
		// 	CronExpression: "30 9-24 * * *", // 9点-24点,1小时执行一次
		// 	ID:             "舔狗日记",
		// })
	})
}
