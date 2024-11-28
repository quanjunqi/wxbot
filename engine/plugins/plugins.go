package plugins

import (
	// 定时任务插件
	_ "wxbot/plugins/cronjob"
	//chatgpt
	_ "wxbot/plugins/chatgpt"
	//chatroom
	_ "wxbot/plugins/chatroom"
	//duanju
	_ "wxbot/plugins/duanju"
	//wanlianli
	_ "wxbot/plugins/wanlianli"
	//fuli
	_ "wxbot/plugins/Fuli"
	//alipay
	_ "wxbot/plugins/alipay"
	//yuying
	_ "wxbot/plugins/yuying"
	//faceold
	_ "wxbot/plugins/faceold"
)
