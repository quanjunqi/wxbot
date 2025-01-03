package bot

import "regexp"

// RegexRule 检查消息是否匹配正则表达式
func RegexRule(regexPattern string) Rule {
	regex := regexp.MustCompile(regexPattern)
	return func(ctx *Ctx) bool {
		if !ctx.IsText() {
			return false
		}
		msg := ctx.Event.Message.Content
		if matched := regex.FindStringSubmatch(msg); matched != nil {
			ctx.State["regex_matched"] = matched
			return true
		}
		return false
	}
}

// OnlyAtMe 只允许@机器人使用，注意这里私聊也是返回true，如仅需群聊，请再加一个OnlyGroup规则
func OnlyAtMe(ctx *Ctx) bool {
	return ctx.IsAt()
}

// OnlyGroup 只允许群聊使用
func OnlyGroup(ctx *Ctx) bool {
	return ctx.IsEventGroupChat()
}

// FullMatchRule 检查消息是否完全匹配
func FullMatchRule(src ...string) Rule {
	return func(ctx *Ctx) bool {
		if !ctx.IsText() {
			return false
		}
		msg := ctx.MessageString()
		for _, str := range src {
			if str == msg {
				ctx.State["matched"] = msg
				return true
			}
		}
		return false
	}
}
