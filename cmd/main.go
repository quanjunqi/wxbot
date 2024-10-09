package main

import (
	"log"
	"time"
	"wxbot/engine/bot"
	"wxbot/engine/pkg/net"

	_ "wxbot/engine/plugins" //插件

	"wxbot/weferry"

	"github.com/spf13/viper"
)

// 初始化解析配置文件，并实例对象
func main() {
	v := viper.New()
	v.SetConfigFile("../config.yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("[main] 读取配置文件失败: %s", err.Error())
	}
	c := bot.NewConfig()
	if err := v.Unmarshal(c); err != nil {
		log.Fatalf("[main] 解析配置文件失败: %s", err.Error())
	}

	f := bot.IFramework(nil)
	f = bot.IFramework(weferry.New(c.BootWxID, c.HookApiUrl))

	// 根据配置文件里的url 检查hook 框架是否能ping通
	if ipPort, err := net.CheckoutIpPort(c.HookApiUrl); err == nil {
		if ping := net.PingConn(ipPort, time.Second*10); !ping {
			c.SetConnHookStatus(false)
			log.Fatalf("[main] 无法连接wefrry框架，网络无法Ping通，请检查网络")
		}
	}
	bot.Run(c, f)
}
