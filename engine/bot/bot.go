package bot

import (
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"
	"wxbot/engine/pkg/log"
)

var (
	bot         *Bot         // 当前机器人
	eventBuffer *EventBuffer // 事件缓冲区
)

type Bot struct {
	config    *Config
	framework IFramework
}

// Run 运行并阻塞主线程，等待事件
func Run(c *Config, f IFramework) {
	if c.BufferLen == 0 {
		c.BufferLen = 4096
	}
	if c.Latency == 0 {
		c.Latency = time.Second
	}
	if c.MaxProcessTime == 0 {
		c.MaxProcessTime = time.Minute * 3
	}
	bot = &Bot{config: c, framework: f}
	log.Printf("[bot] 机器人%s开始工作", c.BootName)
	//启动事件总线
	eventBuffer = NewEventBuffer(bot.config.BufferLen)
	eventBuffer.Loop(bot.config.Latency, bot.config.MaxProcessTime, processEventAsync)

	runServer()
}

// 将当前事件存储到事件上下文中
func processEventAsync(event *Event, framework IFramework, maxWait time.Duration) {
	ctx := &Ctx{
		Bot:       bot,
		State:     State{},
		Event:     event,
		framework: framework,
	}
	matcherLock.Lock()
	if hasMatcherListChanged {
		matcherListForRanging = make([]*Matcher, len(matcherList))
		copy(matcherListForRanging, matcherList)
		hasMatcherListChanged = false
	}
	matcherLock.Unlock()

	preProcessMessageEvent(ctx, event)
	go match(ctx, matcherListForRanging, maxWait) //匹配器开始匹配事件

}

// match 延迟 (1~100ms) 再处理事件
func match(ctx *Ctx, matchers []*Matcher, maxWait time.Duration) {
	goRule := func(rule Rule) <-chan bool {
		ch := make(chan bool, 1)
		go func() {
			defer func() {
				close(ch)
				if err := recover(); err != nil {
					log.Errorf("[robot]执行Rule时运行时发生错误: %v\n%v", err, string(debug.Stack()))
				}
			}()
			ch <- rule(ctx)
		}()
		return ch
	}
	goHandler := func(h Handler) <-chan struct{} {
		ch := make(chan struct{}, 1)
		go func() {
			defer func() {
				close(ch)
				if err := recover(); err != nil {
					log.Errorf("[robot]执行Handler时运行时发生错误: %v\n%v", err, string(debug.Stack()))
				}
			}()
			h(ctx)
			ch <- struct{}{}
		}()
		return ch
	}
	time.Sleep(time.Duration(rand.Intn(500)+1) * time.Millisecond)
	t := time.NewTimer(maxWait)
	defer t.Stop()
loop:
	for _, matcher := range matchers {
		for k := range ctx.State {
			delete(ctx.State, k)
		}
		m := matcher.copy()
		ctx.matcher = m
		// 处理rule
		for _, rule := range m.Rules {
			c := goRule(rule)
			for {
				select {
				case ok := <-c:
					if !ok {
						if m.Break {
							break loop
						}
						continue loop
					}
				case <-t.C:
					if m.NoTimeout {
						t.Reset(maxWait)
						continue
					}
					break loop
				}
				break
			}
		}
		// 处理handler
		if m.Handler != nil {
			c := goHandler(m.Handler)
			for {
				select {
				case <-c:
				case <-t.C:
					if m.NoTimeout {
						t.Reset(maxWait)
						continue
					}
					break loop
				}
				break
			}
		}
		if m.Block {
			break loop
		}
	}
}

// 事件预处理

func preProcessMessageEvent(ctx *Ctx, e *Event) {
	switch e.Type {
	case EventPrivateChat:
		if ctx.IsReference() {
			log.Println(fmt.Sprintf("[回调]收到私聊引用消息(%s[%s])文本消息 ==> %v", e.FromWxId, e.FromName, e.Message.Content))
		}
		if ctx.IsText() {
			log.Println(fmt.Sprintf("[回调]收到私聊(%s[%s])文本消息 ==> %v", e.FromWxId, e.FromName, e.Message.Content))
		}
		if ctx.IsImage() {
			log.Println(fmt.Sprintf("[回调]收到私聊(%s[%s])图片消息 ==> %v", e.FromWxId, e.FromName, e.Message.Id))
		}
	case EventGroupChat:
		if ctx.IsReference() {
			log.Println(fmt.Sprintf("[回调]收到群聊引用消息(%s[%s])>用户(%s[%s])文本消息 ==> %v", e.FromGroupName, e.FromGroup, e.FromWxId, e.FromName, e.Message.Content))
		}
		if ctx.IsText() {
			log.Println(fmt.Sprintf("[回调]收到群聊(%s[%s])>用户(%s[%s])文本消息 ==> %v", e.FromGroupName, e.FromGroup, e.FromWxId, e.FromName, e.Message.Content))
		}
		if ctx.IsImage() {
			log.Println(fmt.Sprintf("[回调]收到私聊(%s[%s])图片消息 ==> %v", e.FromWxId, e.FromName, e.Message.Id))
		}
		if ctx.Ispat() {
			log.Println(fmt.Sprintf("[回调]收到群聊(%s[%s])怕一拍消息 ==> %v", e.FromGroupName, e.FromGroup, e.Message.Content))
		}
	}
}
