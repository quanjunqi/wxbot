package bot

import (
	"sort"
	"sync"
)

// 事件过滤器
type (
	Handler func(ctx *Ctx)      // 事件处理
	Rule    func(ctx *Ctx) bool // 用于过滤事件
)

var (
	matcherList           = make([]*Matcher, 0) // 所有主匹配器列表
	matcherLock           = sync.RWMutex{}      // Matcher 修改读写锁
	matcherListForRanging []*Matcher            // 用于迭代的所有主匹配器列表
	hasMatcherListChanged bool                  // 是否 MatcherList 已经改变，如果改变，下次迭代需要更新
)

// Matcher 是匹配和处理事件的最小单元
type Matcher struct {

	// Block 是否阻断后续 Matcher，为true时当前Matcher匹配成功后，后续Matcher不参与匹配
	Block bool
	// Break 是否退出后续匹配流程, 只有 Rule 返回false且此值为真才会退出, 且不对 mid handler以下的 Rule 生效
	Break bool
	// Priority 优先级，越小优先级越高
	Priority uint64
	// Rules 匹配规则
	Rules []Rule
	// Handler 处理事件的函数
	Handler Handler
	// NoTimeout 处理是否不设超时
	NoTimeout bool
	// Engine 注册 Matcher 的 Engine，Engine可为一系列 Matcher 添加通用 Rule 和其他钩子
	Engine *Engine
}

// State 用来存储匹配器的上下文
type State map[string]interface{}

// 按优先级排序
func sortMatcher() {
	sort.Slice(matcherList, func(i, j int) bool {
		return matcherList[i].Priority < matcherList[j].Priority
	})
	hasMatcherListChanged = true
}

// StoreMatcher 向匹配器列表中添加一个匹配器
func StoreMatcher(m *Matcher) *Matcher {
	matcherLock.Lock()
	defer matcherLock.Unlock()
	if m.Engine != nil {
		m.Block = m.Block || m.Engine.block
	}
	matcherList = append(matcherList, m)
	sortMatcher()
	return m
}

// SetBlock 设置是否阻断后面的 Matcher 触发
func (m *Matcher) SetBlock(block bool) *Matcher {
	m.Block = block
	return m
}

// SetPriority 设置当前 Matcher 优先级
func (m *Matcher) SetPriority(priority uint64) *Matcher {
	matcherLock.Lock()
	defer matcherLock.Unlock()
	m.Priority = priority
	sortMatcher()
	return m
}

// Delete 从匹配器列表中删除当前匹配器
func (m *Matcher) Delete() {
	matcherLock.Lock()
	defer matcherLock.Unlock()
	for i, matcher := range matcherList {
		if m == matcher {
			matcherList = append(matcherList[:i], matcherList[i+1:]...)
			hasMatcherListChanged = true
		}
	}
}
func (m *Matcher) copy() *Matcher {
	return &Matcher{
		Rules:    m.Rules,
		Block:    m.Block,
		Priority: m.Priority,
		Handler:  m.Handler,
		Engine:   m.Engine,
	}
}

// Handle 直接处理事件
func (m *Matcher) Handle(handler Handler) *Matcher {
	m.Handler = handler
	return m
}
