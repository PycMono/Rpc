package model

// 中心服务器响应结果状态对象
type ResultStatus struct {
	// 状态值(成功是0，非成功以负数来表示)
	Code int

	// 英文信息
	Message string

	// 中文描述
	Desc string `json:"-"`
}

func newResultStatus(code int, message, desc string) *ResultStatus {
	return &ResultStatus{
		Code:    code,
		Message: message,
		Desc:    desc,
	}
}

// 定义所有的响应结果的状态枚举值
var (
	Success = newResultStatus(0, "Success", "成功")

	DataError             = newResultStatus(-1, "DataError", "数据错误")
	ClientDataError       = newResultStatus(-2, "ClientDataError", "客户端数据错误")
	SignError             = newResultStatus(-3, "SignError", "签名错误")
	ChannelTypeNotDefined = newResultStatus(-4, "ChannelTypeNotDefined", "聊天频道未定义")
	NoTargetMethod        = newResultStatus(-5, "NoTargetMethod", "找不到目标方法")
	ParamNotMatch         = newResultStatus(-6, "ParamNotMatch", "参数不匹配")
	ParamInValid          = newResultStatus(-7, "ParamInValid", "参数无效")

	NoLogin              = newResultStatus(-11, "NoLogin", "尚未登陆")
	NotInUnion           = newResultStatus(-12, "NotInUnion", "不在公会中")
	NotInShimen          = newResultStatus(-13, "NotInShimen", "不在师门中")
	NotFoundTarget       = newResultStatus(-14, "NotFoundTarget", "未找到目标玩家")
	PlayerNotExist       = newResultStatus(-15, "PlayerNotExist", "玩家不存在")
	ServerGroupNotExist  = newResultStatus(-16, "ServerGroupNotExist", "服务器组不存在")
	NotInTeam            = newResultStatus(-17, "NotInTeam", "不在队伍中")
	LoginOnAnotherDevice = newResultStatus(-18, "LoginOnAnotherDevice", "在另一台设备上登录")

	CantSendMessageToSelf = newResultStatus(-21, "CantSendMessageToSelf", "不能给自己发消息")
	InvalidIP             = newResultStatus(-22, "InvalidIP", "IP无效")
	ResourceNotEnough     = newResultStatus(-23, "ResourceNotEnough", "资源不足")
	NetworkError          = newResultStatus(-24, "NetworkError", "网络错误")

	ContainForbiddenWord = newResultStatus(-31, "ContainForbiddenWord", "含有屏蔽词语")
	SendMessageTooFast   = newResultStatus(-32, "SendMessageTooFast", "发送消息太快")
	LvIsNotEnough        = newResultStatus(-33, "LvIsNotEnough", "等级不足，系统未开放")
	RepeatTooMuch        = newResultStatus(-34, "RepeatTooMuch", "重复次数太多")
)
