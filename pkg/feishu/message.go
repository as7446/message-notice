package feishu

// MsgType 定义飞书机器人消息类型
type MsgType string

const (
	MsgTypeText MsgType = "text"
	MsgTypePost MsgType = "post"
)

// signer 用于在发送前注入签名字段
type signer interface {
	applySignature(ts int64, sign string)
}
