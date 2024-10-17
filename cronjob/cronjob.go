package cronjob

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CronjobBuffer struct { //事件总线
	sync.Mutex
	Cronjoitems chan CronjoBufferItem
	done        chan struct{}
	tasks       map[string]bool // 存储已添加的任务
	mu          sync.Mutex      // 保护任务的并发访问
}

type CronjoBufferItem struct { //事件通道
	ID             string // 任务的唯一标识符
	Function       func() // 要执行的函数
	CronExpression string // Cron 表达式
}

func NewCronjobBuffer(bufferSize int) *CronjobBuffer {
	return &CronjobBuffer{
		Cronjoitems: make(chan CronjoBufferItem, bufferSize),
		done:        make(chan struct{}),
		tasks:       make(map[string]bool),
	}
}

func (cron *CronjobBuffer) Loop(latency time.Duration) {
	go func() {
		ticker := time.NewTicker(latency) // 多少秒检查一次
		defer ticker.Stop()

		for {
			select {
			case task := <-cron.Cronjoitems:
				cron.mu.Lock()
				if _, exists := cron.tasks[task.ID]; !exists {
					cron.tasks[task.ID] = true // 标记任务为已存在
					go cron.runTask(task)      // 启动任务的处理
				}
				cron.mu.Unlock()
			case <-ticker.C:
			case <-cron.done:
				return
			}
		}
	}()
}

// 执行定时任务的方法
func (cron *CronjobBuffer) runTask(task CronjoBufferItem) {
	ticker := time.NewTicker(1 * time.Minute) // 每分钟检查一次
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()
			if cron.shouldRun(currentTime, task.CronExpression) {
				task.Function() // 执行任务
			}
		case <-cron.done:
			return
		}
	}
}

// shouldRun 根据 Cron 表达式判断任务是否需要执行
func (cron *CronjobBuffer) shouldRun(t time.Time, cronExp string) bool {
	parts := strings.Split(cronExp, " ")
	if len(parts) != 5 {
		return false // 不正确的 Cron 表达式
	}

	minute := parts[0]
	hour := parts[1]
	day := parts[2]
	month := parts[3]
	week := parts[4]

	// 检查分钟
	if !matchCronField(t.Minute(), minute) {
		return false
	}

	// 检查小时
	if !matchCronField(t.Hour(), hour) {
		return false
	}

	// 检查日
	if !matchCronField(t.Day(), day) {
		return false
	}

	// 检查月
	if !matchCronField(int(t.Month()), month) {
		return false
	}

	// 检查周
	if !matchCronField(int(t.Weekday()), week) {
		return false
	}

	return true
}

// matchCronField 判断时间字段是否匹配 Cron 表达式
func matchCronField(value int, field string) bool {
	if field == "*" {
		return true
	}

	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		if len(parts) == 2 {
			base := 0
			if parts[0] != "" {
				base, _ = strconv.Atoi(parts[0])
			}
			step, _ := strconv.Atoi(parts[1])
			return value >= base && (value-base)%step == 0
		}
	}
	// 处理范围
	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		if len(parts) == 2 {
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			return value >= start && value <= end
		}
	}
	parts := strings.Split(field, ",")
	for _, part := range parts {
		if part == fmt.Sprintf("%d", value) {
			return true
		}
	}
	return false
}

func (cron *CronjobBuffer) AddTask(task CronjoBufferItem) {
	cron.Cronjoitems <- task // 将任务添加到通道
}

func (cron *CronjobBuffer) Stop() {
	close(cron.done)
}
