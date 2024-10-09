package bot

const (
	EventGroupChat           = "EventGroupChat"           // 群聊消息事件
	EventPrivateChat         = "EventPrivateChat"         // 私聊消息事件
	EventMPChat              = "EventMPChat"              // 公众号消息事件
	EventSelfMessage         = "EventSelfMessage"         // 自己发的消息事件
	EventFriendVerify        = "EventFriendVerify"        // 好友请求事件
	EventTransfer            = "EventTransfer"            // 好友转账事件
	EventMessageWithdraw     = "EventMessageWithdraw"     // 消息撤回事件
	EventSystem              = "EventSystem"              // 系统消息事件
	EventGroupMemberIncrease = "EventGroupMemberIncrease" // 群成员增加事件
	EventGroupMemberDecrease = "EventGroupMemberDecrease" // 群成员减少事件
	EventInvitedInGroup      = "EventInvitedInGroup"      // 被邀请入群事件
)
const (
	MsgTypeText           = 1     // 文本消息
	MsgTypeImage          = 3     // 图片消息
	MsgTypeVoice          = 34    // 语音消息
	MsgTypeAuthentication = 37    // 认证消息
	MsgTypePossibleFriend = 40    // 好友推荐消息
	MsgTypeShareCard      = 42    // 名片消息
	MsgTypeVideo          = 43    // 视频消息
	MsgTypeMemePicture    = 47    // 表情消息
	MsgTypeLocation       = 48    // 地理位置消息
	MsgTypeApp            = 49    // APP消息
	MsgTypeMicroVideo     = 62    // 小视频消息
	MsgTypeSystem         = 10000 // 系统消息
	MsgTypeRecalled       = 10002 // 消息撤回
	MsgTypeReference      = 10003 // 消息引用
)

// IsText 判断消息类型是否为文本
func (ctx *Ctx) IsText() bool {
	return ctx.Event.Message != nil && ctx.Event.Message.Type == MsgTypeText
}

// IsAt 判断是否被@了，仅在群聊中有效，私聊也算被@了
func (ctx *Ctx) IsAt() bool {
	return ctx.Event.IsAtMe
}

// IsEventPrivateChat 判断消息是否是私聊消息
func (ctx *Ctx) IsEventPrivateChat() bool {
	return ctx.Event.Type == EventPrivateChat
}
