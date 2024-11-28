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
		priority: o.Priority, // 直接使用全局 priority
		service:  service,
	}
	// o.Priority = priority // 这里可以选择是否需要更新 Options 的 Priority
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

// OnFullMatch 完全匹配触发器
func (e *Engine) OnFullMatch(src string, rules ...bot.Rule) *Matcher {
	return (*Matcher)(e.en.OnFullMatch(src, rules...).SetPriority(e.priority))
}

// OnFullMatchGroup 完全匹配触发器组
func (e *Engine) OnFullMatchGroup(src []string, rules ...bot.Rule) *Matcher {
	return (*Matcher)(e.en.OnFullMatchGroup(src, rules...).SetPriority(e.priority))
}