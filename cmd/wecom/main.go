package main

import (
	"flag"
	"log"

	"github.com/as7446/message-notice/internal/cliutil"
	"github.com/as7446/message-notice/pkg/wecom"
)

func main() {
	var webhook string
	var msgType string
	var text string
	var atUserIDs string
	var atMobiles string

	flag.StringVar(&webhook, "webhook", "", "企业微信机器人 webhook，形如 https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx")
	flag.StringVar(&msgType, "type", "text", "消息类型：text|markdown")
	flag.StringVar(&text, "text", "", "文本或 Markdown 内容")
	flag.StringVar(&atUserIDs, "at-userids", "", "以逗号分隔的 userid 列表，用于 @ 提醒")
	flag.StringVar(&atMobiles, "at-mobiles", "", "以逗号分隔的手机号列表，用于 @ 提醒")
	flag.Parse()

	if webhook == "" || text == "" {
		log.Fatal("webhook 与 text 为必填")
	}

	client := wecom.NewClient(webhook)
	userIDs := cliutil.SplitCSV(atUserIDs)
	mobiles := cliutil.SplitCSV(atMobiles)

	switch msgType {
	case "text":
		msg := wecom.NewTextMessage().
			SetText(text).
			SetMentionUserIDs(userIDs).
			SetMentionMobiles(mobiles).
			AppendVisibleMentions(userIDs)
		if _, err := client.Send(msg); err != nil {
			log.Fatal(err)
		}
	case "markdown":
		msg := wecom.NewMarkdownMessage().
			SetContent(text).
			AppendMentions(userIDs)
		if _, err := client.Send(msg); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("不支持的 type: %s", msgType)
	}
}
