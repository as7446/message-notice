package wecom

import (
	"encoding/json"
	"strings"
)

// TextMessage 企业微信文本消息
type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	Text    Text    `json:"text"`
}

// Text 文本内容
type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

// NewTextMessage 创建文本消息
func NewTextMessage() *TextMessage {
	return &TextMessage{MsgType: MsgTypeText}
}

func (m *TextMessage) SetText(text string) *TextMessage {
	m.Text.Content = text
	return m
}

// SetMentionUserIDs 设置按 userid @
func (m *TextMessage) SetMentionUserIDs(userIDs []string) *TextMessage {
	m.Text.MentionedList = dedupStrings(userIDs)
	return m
}

// SetMentionMobiles 设置按手机号 @
func (m *TextMessage) SetMentionMobiles(mobiles []string) *TextMessage {
	m.Text.MentionedMobileList = dedupStrings(mobiles)
	return m
}

func (m *TextMessage) ToBytes() ([]byte, error) {
	return json.Marshal(m)
}

// 将用户标记追加到正文，便于可视化提醒
func (m *TextMessage) AppendVisibleMentions(userIDs []string) *TextMessage {
	tokens := buildUserTokens(userIDs)
	if len(tokens) == 0 {
		return m
	}
	if m.Text.Content == "" {
		m.Text.Content = strings.Join(tokens, " ")
	} else {
		m.Text.Content = strings.TrimSpace(m.Text.Content + "\n" + strings.Join(tokens, " "))
	}
	return m
}
