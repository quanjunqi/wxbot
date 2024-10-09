package bot

import "time"

type Config struct {
	BootWxID       string        `mapstructure:"bootWxID"`       // 机器人微信ID
	BootName       string        `mapstructure:"bootName"`       // 机器人名字
	HookApiUrl     string        `mapstructure:"hookApiUrl"`     // 接入框架API地址
	ConnHookStatus bool          `mapstructure:"connHookStatus"` // 连接Hook框架状态
	BufferLen      uint          `mapstructure:"-"`              // 事件缓冲区长度, 默认4096
	Latency        time.Duration `mapstructure:"-"`              // 事件处理延迟 (延迟 latency + (0~100ms) 再处理事件) (默认1s)
	MaxProcessTime time.Duration `mapstructure:"-"`              // 事件最大处理时间 (默认3min)
}

func NewConfig() *Config {
	return &Config{
		ConnHookStatus: true,
	}
}

// SetConnHookStatus 设置连接Hook框架状态
func (c *Config) SetConnHookStatus(status bool) {
	c.ConnHookStatus = status
}
