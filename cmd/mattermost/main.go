package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/as7446/message-notice/pkg/mattermost"
)

func main() {
	var webhook string
	var text string
	var channel string
	var username string
	var iconURL string
	var iconEmoji string
	var attach string // 简易格式：title|text|color
	var attachmentsJSON string
	var attachmentsFile string
	var mentions string // 以英文逗号分隔的用户名，将在文本中可见@

	flag.StringVar(&webhook, "webhook", "", "Mattermost 入站 webhook URL")
	flag.StringVar(&text, "text", "", "消息内容")
	flag.StringVar(&channel, "channel", "", "目标频道名称或 ID")
	flag.StringVar(&username, "username", "", "覆盖用户名显示")
	flag.StringVar(&iconURL, "icon-url", "", "覆盖头像 URL")
	flag.StringVar(&iconEmoji, "icon-emoji", "", "覆盖头像 Emoji，例如 :rocket:")
	flag.StringVar(&attach, "attach", "", "'title|text|color' 附件简易格式；可多次运行附加")
	flag.StringVar(&attachmentsJSON, "attachments-json", "", "完整附件的 JSON 字符串（单对象或数组）")
	flag.StringVar(&attachmentsFile, "attachments-file", "", "包含附件 JSON 的文件路径（单对象或数组）")
	flag.StringVar(&mentions, "mentions", "", "以逗号分隔的用户名列表用于可见@（例如 user1,user2）")
	flag.Parse()

	if webhook == "" || text == "" {
		log.Fatal("webhook 和 text 为必填")
	}

	// 如果提供了 mentions，则将其前置到文本中
	if mentions != "" {
		parts := strings.Split(mentions, ",")
		var atTokens []string
		for _, m := range parts {
			m = strings.TrimSpace(m)
			if m == "" {
				continue
			}
			if strings.HasPrefix(m, "@") {
				m = m[1:]
			}
			atTokens = append(atTokens, "@"+m)
		}
		if len(atTokens) > 0 {
			text = strings.Join(atTokens, " ") + "\n" + text
		}
	}

	client := mattermost.NewClient(webhook)
	msg := mattermost.NewTextMessage().SetText(text)
	if channel != "" {
		msg.SetChannel(channel)
	}
	if username != "" {
		msg.SetUsername(username)
	}
	if iconURL != "" {
		msg.SetIconUrl(iconURL)
	}
	if iconEmoji != "" {
		msg.SetIconEmoji(iconEmoji)
	}
	if attach != "" {
		parts := strings.SplitN(attach, "|", 3)
		a := mattermost.Attachment{}
		if len(parts) > 0 {
			a.Title = parts[0]
		}
		if len(parts) > 1 {
			a.Text = parts[1]
		}
		if len(parts) > 2 {
			a.Color = parts[2]
		}
		msg.AddAttachment(a)
	}

	// 从 JSON 字符串解析附件
	if attachmentsJSON != "" {
		addAttachmentsFromJSON(msg, attachmentsJSON)
	}
	// 从 JSON 文件解析附件
	if attachmentsFile != "" {
		b, err := os.ReadFile(attachmentsFile)
		if err != nil {
			log.Fatal(err)
		}
		addAttachmentsFromJSON(msg, string(b))
	}

	if _, err := client.Send(msg); err != nil {
		log.Fatal(err)
	}
}

// addAttachmentsFromJSON 支持单对象或数组形式的附件
func addAttachmentsFromJSON(msg *mattermost.TextMessage, jsonStr string) {
	jsonStr = strings.TrimSpace(jsonStr)
	if jsonStr == "" {
		return
	}
	// 优先尝试数组
	var arr []mattermost.Attachment
	if err := json.Unmarshal([]byte(jsonStr), &arr); err == nil {
		for _, a := range arr {
			msg.AddAttachment(a)
		}
		return
	}
	// 尝试单对象
	var one mattermost.Attachment
	if err := json.Unmarshal([]byte(jsonStr), &one); err == nil {
		msg.AddAttachment(one)
	}
}
