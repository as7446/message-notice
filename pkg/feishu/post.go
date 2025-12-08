package feishu

import (
	"encoding/json"
	"strconv"
	"strings"
)

// PostMessage 结构化富文本（post）消息
// 为简化使用，这里提供将多行文本转换为段落的构造器
type PostMessage struct {
	MsgType   MsgType  `json:"msg_type"`
	Content   PostBody `json:"content"`
	Timestamp string   `json:"timestamp,omitempty"`
	Sign      string   `json:"sign,omitempty"`
}

// PostBody 封装多语言内容；目前只覆盖 zh_cn
type PostBody struct {
	Post PostLangWrapper `json:"post"`
}

type PostLangWrapper struct {
	ZhCN PostLang `json:"zh_cn"`
}

type PostLang struct {
	Title   string          `json:"title"`
	Content [][]PostElement `json:"content"`
}

// PostElement 与官方字段保持一致，当前仅使用 text/link 两种
type PostElement struct {
	Tag  string `json:"tag"`
	Text string `json:"text,omitempty"`
	Href string `json:"href,omitempty"`
}

// NewPostMessage 创建 post 消息
func NewPostMessage() *PostMessage {
	return &PostMessage{
		MsgType: MsgTypePost,
		Content: PostBody{
			Post: PostLangWrapper{
				ZhCN: PostLang{Content: [][]PostElement{}},
			},
		},
	}
}

// SetTitle 设置卡片标题
func (m *PostMessage) SetTitle(title string) *PostMessage {
	m.Content.Post.ZhCN.Title = title
	return m
}

// SetTextLines 将多行文本转换为多段内容，每行一个段落
func (m *PostMessage) SetTextLines(lines []string) *PostMessage {
	var content [][]PostElement
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		content = append(content, []PostElement{{Tag: "text", Text: line}})
	}
	m.Content.Post.ZhCN.Content = content
	return m
}

// ToBytes 序列化
func (m *PostMessage) ToBytes() ([]byte, error) {
	return json.Marshal(m)
}

// applySignature 注入签名
func (m *PostMessage) applySignature(ts int64, sign string) {
	m.Timestamp = strconv.FormatInt(ts, 10)
	m.Sign = sign
}
