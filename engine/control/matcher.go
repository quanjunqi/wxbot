package control

import "wxbot/engine/bot"

type Matcher bot.Matcher

// SetBlock 设置是否阻断后面的Matcher触发
func (m *Matcher) SetBlock(block bool) *Matcher {
	_ = (*bot.Matcher)(m).SetBlock(block)
	return m
}

// SetPriority 设置当前Matcher优先级
func (m *Matcher) SetPriority(priority uint64) *Matcher {
	_ = (*bot.Matcher)(m).SetPriority(priority)
	return m
}

// Handle 直接处理事件
func (m *Matcher) Handle(handler bot.Handler) {
	_ = (*bot.Matcher)(m).Handle(handler)
}
