package control

import (
	"log"
	"sync/atomic"
)

var (
	priority    uint64
	priorityMap = make(map[uint64]string)
)

// Register 注册插件控制器
func Register(service string, o *Options) *Engine {
	atomic.AddUint64(&priority, 10)
	s, ok := priorityMap[priority]
	if ok {
		log.Fatalf("插件[%s]优先级 %d 已被 %s 占用", service, priority, s)
	}
	priorityMap[priority] = service
	log.Printf("插件[%s]已注册, 优先级: %d", service, o.Priority)
	return newEngine(service, o)
}
