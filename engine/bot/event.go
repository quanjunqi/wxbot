package bot

// Event 记录一次回调事件
type Event struct {
	Type           string   // 消息类型
	RobotWxId      string   // 机器人微信id
	IsAtMe         bool     // 机器人是否被@了，@所有人不算
	FromUniqueID   string   // 消息来源唯一id, 私聊为发送者微信id, 群聊为群id
	FromUniqueName string   // 消息来源唯一名称, 私聊为发送者昵称, 群聊为群名称
	FromWxId       string   // 消息来源微信id
	FromName       string   // 消息来源昵称
	FromGroup      string   // 消息来源群id
	FromGroupName  string   // 消息来源群名称
	RawMessage     string   // 原始消息
	Message        *Message // 消息内容
}

// Message 记录消息的具体内容
type Message struct {
	Id      string // 消息id
	Type    int64  // 消息类型
	Content string // 消息内容
}
