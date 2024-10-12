package cronjob

import (
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
	ID            string    // 任务的唯一标识符
	Function      func()    // 要执行的函数
	ExecutionTime time.Time // 任务的下次执行时间
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

func (cron *CronjobBuffer) runTask(task CronjoBufferItem) {
	ticker := time.NewTicker(1 * time.Minute) // 每分钟检查一次
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()
			if currentTime.Year() == task.ExecutionTime.Year() &&
				currentTime.YearDay() == task.ExecutionTime.YearDay() &&
				currentTime.Hour() == task.ExecutionTime.Hour() &&
				currentTime.Minute() == task.ExecutionTime.Minute() {
				// 更新下一次执行时间为明天
				task.Function() // 执行任务
				task.ExecutionTime = task.ExecutionTime.Add(24 * time.Hour)
			}

		case <-cron.done:
			return
		}
	}
}

func (cron *CronjobBuffer) AddTask(task CronjoBufferItem) {
	cron.Cronjoitems <- task // 将任务添加到通道
}

func (cron *CronjobBuffer) Stop() {
	close(cron.done)
}
