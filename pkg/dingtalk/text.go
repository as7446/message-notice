/*
@Author : Fuxuhao
@Time : 2024/7/26 10:52
*/
package dingtalk

import "encoding/json"

// TextMessage 钉钉机器人 文本消息
// 支持通过手机号或用户ID进行@提醒
// 注意：要在会话里看到可见的@，文本中还需手动写入 @手机号
// （at 字段用于触发提醒，文本中的 @ 用于可视化展示）
type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	Text    Text    `json:"text"`
	At      At      `json:"at"`
}

// Text 文本消息体
// Content 为消息内容
type Text struct {
	Content string `json:"content"`
}

// NewTextMessage 创建 Text 文本消息
func NewTextMessage() *TextMessage {
	return &TextMessage{MsgType: MsgTypeText}
}

// SetText 设置文本内容
func (t *TextMessage) SetText(content string) *TextMessage {
	t.Text = Text{Content: content}
	return t
}

// SetAt 通过手机号@，并设置是否@所有人
func (t *TextMessage) SetAt(atMobiles []string, isAtAll bool) *TextMessage {
	t.At.AtMobiles = atMobiles
	t.At.IsAtAll = isAtAll
	return t
}

// SetAtUserIds 通过用户ID@，并设置是否@所有人
func (t *TextMessage) SetAtUserIds(atUserIds []string, isAtAll bool) *TextMessage {
	t.At.AtUserIds = atUserIds
	t.At.IsAtAll = isAtAll
	return t
}

// ToBytes 序列化
func (t *TextMessage) ToBytes() ([]byte, error) {
	b, err := json.Marshal(t)
	return b, err
}
