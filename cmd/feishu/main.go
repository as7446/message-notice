package main

import (
	"flag"
	"log"
	"strings"

	"github.com/as7446/message-notice/internal/cliutil"
	"github.com/as7446/message-notice/pkg/feishu"
)

func main() {
	var webhook string
	var secret string
	var msgType string
	var title string
	var text string
	var at string

	flag.StringVar(&webhook, "webhook", "", "飞书机器人 Webhook，形如 https://open.feishu.cn/open-apis/bot/v2/hook/xxx")
	flag.StringVar(&secret, "secret", "", "飞书机器人签名密钥（可选）")
	flag.StringVar(&msgType, "type", "text", "消息类型：text|post")
	flag.StringVar(&title, "title", "", "post 消息标题（type=post 时可选）")
	flag.StringVar(&text, "text", "", "消息内容，多行可直接换行输入")
	flag.StringVar(&at, "at", "", "以逗号分隔的 user_id/open_id 列表，用于 @ 提醒（仅 text 支持自动拼接）")
	flag.Parse()

	if webhook == "" || text == "" {
		log.Fatal("webhook 与 text 为必填")
	}

	client := feishu.NewClient(webhook, secret)
	userIDs := cliutil.SplitCSV(at)

	switch msgType {
	case "text":
		msg := feishu.NewTextMessage().SetText(text)
		if len(userIDs) > 0 {
			msg.MentionUserIDs(userIDs)
		}
		if _, err := client.Send(msg); err != nil {
			log.Fatal(err)
		}
	case "post":
		lines := strings.Split(text, "\n")
		msg := feishu.NewPostMessage().
			SetTitle(title).
			SetTextLines(lines)
		if _, err := client.Send(msg); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("不支持的 type: %s", msgType)
	}
}
