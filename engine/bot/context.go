package bot

import (
	"strings"
	"sync"
)

// 事件上下文
type Ctx struct {
	matcher   *Matcher //事件匹配器
	Bot       *Bot
	Event     *Event
	State     State
	framework IFramework

	// lazy message
	once    sync.Once
	mutex   sync.Mutex
	message string
}

// GetMatcher 获取匹配器
func (ctx *Ctx) GetMatcher() *Matcher {
	return ctx.matcher
}

// MessageString 字符串消息便于Regex
func (ctx *Ctx) MessageString() string {
	ctx.once.Do(func() {
		if ctx.Event != nil && ctx.IsText() {
			if !ctx.IsAt() || ctx.IsEventPrivateChat() {
				ctx.message = ctx.Event.Message.Content
			} else {
				ctx.message = strings.TrimPrefix(ctx.Event.Message.Content, "@"+bot.config.BootName)
				ctx.message = strings.ReplaceAll(ctx.message, "\u2005", "")
				ctx.message = strings.TrimSpace(ctx.message)
			}
		}
	})
	return ctx.message
}
