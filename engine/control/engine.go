package control

import (
	"wxbot/engine/bot"

)

type Engine struct {
	en       *bot.Engine // robot engine
	priority uint64      // 优先级
	service  string      // 插件服务名
}

func newEngine(service string, o *Options) (e *Engine) {
	e = &Engine{
		en:       bot.New(),
		priority: priority,
		service:  service,
	}
	o.Priority = priority
	return
}

// OnMessage 消息触发器
func (e *Engine) OnMessage(rules ...bot.Rule) *Matcher {
	return (*Matcher)(e.en.On(rules...).SetPriority(e.priority))
}

// OnRegex 正则触发器
func (e *Engine) OnRegex(regexPattern string, rules ...bot.Rule) *Matcher {
	return (*Matcher)(e.en.OnRegex(regexPattern, rules...).SetPriority(e.priority))
}
