package feishu

import (
	"encoding/json"
	"strconv"
	"strings"
)

// TextMessage 文本消息；支持在内容中自动插入用户 @ 标记
type TextMessage struct {
	MsgType   MsgType     `json:"msg_type"`
	Content   TextContent `json:"content"`
	Timestamp string      `json:"timestamp,omitempty"`
	Sign      string      `json:"sign,omitempty"`
	extras    []string    // 额外拼接在文本末尾的 @ 标记
}

// TextContent 文本内容
type TextContent struct {
	Text string `json:"text"`
}

// NewTextMessage 创建文本消息
func NewTextMessage() *TextMessage {
	return &TextMessage{MsgType: MsgTypeText}
}

// SetText 设置文本内容
func (m *TextMessage) SetText(text string) *TextMessage {
	m.Content.Text = text
	return m
}

// MentionUserIDs 在文本末尾追加 @user_id 标记
// 飞书需要在文本中手动写入 <at user_id="xxx">姓名</at> 才能可见提醒，这里自动生成占位。
func (m *TextMessage) MentionUserIDs(userIDs []string) *TextMessage {
	m.extras = append(m.extras, buildAtTokens(userIDs)...)
	return m
}

// MentionOpenIDs 与 MentionUserIDs 等价，命名方便调用
func (m *TextMessage) MentionOpenIDs(openIDs []string) *TextMessage {
	return m.MentionUserIDs(openIDs)
}

// ToBytes 序列化为请求体
func (m *TextMessage) ToBytes() ([]byte, error) {
	if len(m.extras) > 0 {
		if m.Content.Text == "" {
			m.Content.Text = strings.Join(m.extras, " ")
		} else {
			m.Content.Text = strings.TrimSpace(m.Content.Text + "\n" + strings.Join(m.extras, " "))
		}
	}
	return json.Marshal(m)
}

// applySignature 注入签名字段（客户端在发送前调用）
func (m *TextMessage) applySignature(ts int64, sign string) {
	m.Timestamp = strconv.FormatInt(ts, 10)
	m.Sign = sign
}

// buildAtTokens 构造 <at user_id="..."> 占位
func buildAtTokens(userIDs []string) []string {
	var tokens []string
	for _, id := range userIDs {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		tokens = append(tokens, `<at user_id="`+id+`">`+id+`</at>`)
	}
	return tokens
}
