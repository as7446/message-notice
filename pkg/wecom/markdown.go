package wecom

import (
	"encoding/json"
	"strings"
)

// MarkdownMessage 企业微信 Markdown 消息
// @ 提醒需要在内容中显式写入 <@userid>
type MarkdownMessage struct {
	MsgType  MsgType  `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}

func NewMarkdownMessage() *MarkdownMessage {
	return &MarkdownMessage{MsgType: MsgTypeMarkdown}
}

func (m *MarkdownMessage) SetContent(content string) *MarkdownMessage {
	m.Markdown.Content = content
	return m
}

// AppendMentions 在内容末尾增加 <@userid> 便于可视化提醒
func (m *MarkdownMessage) AppendMentions(userIDs []string) *MarkdownMessage {
	if len(userIDs) == 0 {
		return m
	}
	tokens := buildUserTokens(userIDs)
	if len(tokens) == 0 {
		return m
	}
	if m.Markdown.Content == "" {
		m.Markdown.Content = strings.Join(tokens, " ")
	} else {
		m.Markdown.Content = strings.TrimSpace(m.Markdown.Content + "\n" + strings.Join(tokens, " "))
	}
	return m
}

func (m *MarkdownMessage) ToBytes() ([]byte, error) {
	return json.Marshal(m)
}
