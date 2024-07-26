/*
@Author : Fuxuhao
@Time : 2024/7/26 18:57
*/
package mattermost

import (
	"encoding/json"
)

type TextMessage struct {
	Text string `json:"text"`
	//Channel     string `json:"channel"`
	//Username    string `json:"username"`
	//IconUrl     string `json:"icon_url"`
	//IconEmoji   string `json:"icon_emoji"`
	//Attachments string `json:"attachments"`
	//Type        string `json:"type"`
	//Props       string `json:"props"`
	//Priority    string `json:"priority"`
}

func NewTextMessage() *TextMessage {
	return &TextMessage{}
}

//	func (t *TextMessage) SetChannel(channel string) *TextMessage {
//		t.Channel = channel
//		return t
//	}
//
//	func (t *TextMessage) SetUsername(username string) *TextMessage {
//		t.Username = username
//		return t
//	}
//
//	func (t *TextMessage) SetIconUrl(url string) *TextMessage {
//		t.IconUrl = url
//		return t
//	}
func (t *TextMessage) SetText(text string) *TextMessage {
	t.Text = text
	return t
}

func (t *TextMessage) ToBytes() ([]byte, error) {
	msg, err := json.Marshal(t)
	return msg, err
}
