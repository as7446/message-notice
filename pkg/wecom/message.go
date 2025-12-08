package wecom

// MsgType 企业微信机器人消息类型
type MsgType string

const (
	MsgTypeText     MsgType = "text"
	MsgTypeMarkdown MsgType = "markdown"
)
