package bot

type Engine struct {
	block    bool       // 是否阻断后续处理器
	matchers []*Matcher // 匹配器
}

// New 生成空引擎
func New() *Engine {
	return &Engine{}
}

var defaultEngine = New()



// SetBlock 设置是否阻断后续处理器
func (e *Engine) SetBlock(block bool) *Engine {
	e.block = block
	return e
}

// 方法重载

// On 添加新的匹配器
func On(rules ...Rule) *Matcher { return defaultEngine.On(rules...) }

// On 添加新的匹配器
func (e *Engine) On(rules ...Rule) *Matcher {
	matcher := &Matcher{
		Engine: e,
		Rules:  rules,
	}
	e.matchers = append(e.matchers, matcher)
	return StoreMatcher(matcher)
}

// OnMessage 消息触发器
func OnMessage(rules ...Rule) *Matcher { return On(rules...) }

// OnMessage 消息触发器
func (e *Engine) OnMessage(rules ...Rule) *Matcher { return e.On(rules...) }

// OnRegex 正则触发器
func OnRegex(regexPattern string, rules ...Rule) *Matcher { //全局函数
	return defaultEngine.OnRegex(regexPattern, rules...)
}

// OnRegex 正则触发器
func (e *Engine) OnRegex(regexPattern string, rules ...Rule) *Matcher { //局部函数
	matcher := &Matcher{
		Engine: e,
		Rules:  append([]Rule{RegexRule(regexPattern)}, rules...),
	}
	e.matchers = append(e.matchers, matcher)
	return StoreMatcher(matcher)
}
