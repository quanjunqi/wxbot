package control

type Options struct {
	Alias    string // 插件别名
	Help     string // 插件帮助信息
	Priority uint64 // 优先级,只读
}
