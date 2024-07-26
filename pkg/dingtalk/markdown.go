/*
@Author : Fuxuhao
@Time : 2024/7/26 10:52
*/
package dingtalk

import "encoding/json"

type MarkdownMessage struct {
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
	MsgType  MsgType  `json:"msgtype"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

func NewMarkdownMessage() *MarkdownMessage {
	msg := &MarkdownMessage{MsgType: MsgTypeMarkdown}
	return msg
}

func (msg *MarkdownMessage) SetMarkdown(title, text string) *MarkdownMessage {
	msg.Markdown = Markdown{
		Title: title,
		Text:  text,
	}
	return msg
}

func (msg *MarkdownMessage) SetAt(atMobiles []string, isAtAll bool) *MarkdownMessage {
	msg.At = At{
		AtMobiles: atMobiles,
		IsAtAll:   isAtAll,
	}
	return msg
}

func (msg *MarkdownMessage) ToBytes() ([]byte, error) {
	b, err := json.Marshal(msg)
	return b, err
}
