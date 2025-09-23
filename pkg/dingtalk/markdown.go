/*
@Author : Fuxuhao
@Time : 2024/7/26 10:52
*/
package dingtalk

import "encoding/json"

// MarkdownMessage 钉钉机器人 Markdown 消息
// 支持通过手机号或用户ID进行@提醒
// 注意：要在会话里看到可见的@，Markdown 文本中还需手动写入 @手机号
// （at 字段用于触发提醒，文本中的 @ 用于可视化展示）
type MarkdownMessage struct {
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
	MsgType  MsgType  `json:"msgtype"`
}

// Markdown Markdown 消息体
// Title 为卡片标题，Text 为 markdown 正文
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// At @提醒配置
// AtMobiles: 按手机号@；AtUserIds: 按用户ID@；IsAtAll: 是否@所有人
type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
	IsAtAll   bool     `json:"isAtAll"`
}

// NewMarkdownMessage 创建 Markdown 消息
func NewMarkdownMessage() *MarkdownMessage {
	msg := &MarkdownMessage{MsgType: MsgTypeMarkdown}
	return msg
}

// SetMarkdown 设置 Markdown 标题与正文
func (msg *MarkdownMessage) SetMarkdown(title, text string) *MarkdownMessage {
	msg.Markdown = Markdown{
		Title: title,
		Text:  text,
	}
	return msg
}

// SetAt 通过手机号@，并设置是否@所有人
func (msg *MarkdownMessage) SetAt(atMobiles []string, isAtAll bool) *MarkdownMessage {
	msg.At.AtMobiles = atMobiles
	msg.At.IsAtAll = isAtAll
	return msg
}

// SetAtUserIds 通过用户ID@，并设置是否@所有人
func (msg *MarkdownMessage) SetAtUserIds(atUserIds []string, isAtAll bool) *MarkdownMessage {
	msg.At.AtUserIds = atUserIds
	msg.At.IsAtAll = isAtAll
	return msg
}

// ToBytes 序列化
func (msg *MarkdownMessage) ToBytes() ([]byte, error) {
	b, err := json.Marshal(msg)
	return b, err
}
