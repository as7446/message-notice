/*
@Author : Fuxuhao
@Time : 2024/7/26 18:57
*/
package mattermost

import (
	"encoding/json"
)

// TextMessage represents Mattermost incoming webhook payload
// https://developers.mattermost.com/integrate/webhooks/incoming/
type TextMessage struct {
	Text        string         `json:"text"`
	Channel     string         `json:"channel,omitempty"`
	Username    string         `json:"username,omitempty"`
	IconUrl     string         `json:"icon_url,omitempty"`
	IconEmoji   string         `json:"icon_emoji,omitempty"`
	Attachments []Attachment   `json:"attachments,omitempty"`
	Props       map[string]any `json:"props,omitempty"`
}

type Attachment struct {
	Fallback   string            `json:"fallback,omitempty"`
	Color      string            `json:"color,omitempty"`
	Pretext    string            `json:"pretext,omitempty"`
	AuthorName string            `json:"author_name,omitempty"`
	AuthorLink string            `json:"author_link,omitempty"`
	AuthorIcon string            `json:"author_icon,omitempty"`
	Title      string            `json:"title,omitempty"`
	TitleLink  string            `json:"title_link,omitempty"`
	Text       string            `json:"text,omitempty"`
	Fields     []AttachmentField `json:"fields,omitempty"`
	ImageUrl   string            `json:"image_url,omitempty"`
	ThumbUrl   string            `json:"thumb_url,omitempty"`
	Footer     string            `json:"footer,omitempty"`
	FooterIcon string            `json:"footer_icon,omitempty"`
	Ts         int64             `json:"ts,omitempty"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func NewTextMessage() *TextMessage {
	return &TextMessage{Props: map[string]any{}}
}

func (t *TextMessage) SetText(text string) *TextMessage {
	t.Text = text
	return t
}

func (t *TextMessage) SetChannel(channel string) *TextMessage {
	t.Channel = channel
	return t
}

func (t *TextMessage) SetUsername(username string) *TextMessage {
	t.Username = username
	return t
}

func (t *TextMessage) SetIconUrl(url string) *TextMessage {
	t.IconUrl = url
	return t
}

func (t *TextMessage) SetIconEmoji(emoji string) *TextMessage {
	t.IconEmoji = emoji
	return t
}

func (t *TextMessage) AddAttachment(a Attachment) *TextMessage {
	t.Attachments = append(t.Attachments, a)
	return t
}

func (t *TextMessage) PutProp(key string, value any) *TextMessage {
	if t.Props == nil {
		t.Props = map[string]any{}
	}
	t.Props[key] = value
	return t
}

func (t *TextMessage) ToBytes() ([]byte, error) {
	msg, err := json.Marshal(t)
	return msg, err
}
