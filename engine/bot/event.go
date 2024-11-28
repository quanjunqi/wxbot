package bot

// Event 记录一次回调事件
type Event struct {
	Type             string            // 消息类型
	RobotWxId        string            // 机器人微信id
	IsAtMe           bool              // 机器人是否被@了，@所有人不算
	FromUniqueID     string            // 消息来源唯一id, 私聊为发送者微信id, 群聊为群id
	FromUniqueName   string            // 消息来源唯一名称, 私聊为发送者昵称, 群聊为群名称
	FromWxId         string            // 消息来源微信id
	FromName         string            // 消息来源昵称
	FromGroup        string            // 消息来源群id
	FromGroupName    string            // 消息来源群名称
	RawMessage       string            // 原始消息
	Message          *Message          // 消息内容
	ReferenceMessage *ReferenceMessage // 引用消息
}

// Message 记录消息的具体内容
type Message struct {
	Id      int64  // 消息id
	Type    int64  // 消息类型
	Content string // 消息内容
}

// ReferenceMessage 记录引用消息的具体内容
type ReferenceMessage struct {
	ReferenceMessageType int64  //引用的消息类型
	ReferenceMessageID   string //引用的消息ID
	FromUser             string // 消息来源群ID
	ChatUser             string // 消息来源微信ID
	Content              string // 消息内容
	ImageURL             string // 引用图片地址

}

// TransferMessage 记录转账消息的具体内容
type TransferMessage struct {
	FromWxId     string // 发送者微信ID
	MsgSource    int64  // 消息来源 1:收到转账 2:对方接收转账 3:发出转账 4:自己接收转账 5:对方退还 6:自己退还
	TransferType int64  // 转账类型 1:即时到账 2:延时到账
	Money        string // 转账金额，单位元
	Memo         string // 转账备注
	TransferId   string // 转账ID
	TransferTime string // 转账时间，10位时间戳
}
