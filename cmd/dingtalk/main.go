/*
@Author : Fuxuhao
@Time : 2024/7/26 11:03
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/as7446/message-notice/pkg/dingtalk"
)

func main() {
	var accessToken string
	var secret string
	var msgType string
	var title string
	var text string
	var atMobiles string
	var atAll bool

	flag.StringVar(&accessToken, "token", "", "钉钉机器人 access token")
	flag.StringVar(&secret, "secret", "", "钉钉机器人 secret（可选）")
	flag.StringVar(&msgType, "type", "text", "消息类型：text|markdown")
	flag.StringVar(&title, "title", "", "Markdown 标题（当 type=markdown 时必填）")
	flag.StringVar(&text, "text", "", "消息内容")
	flag.StringVar(&atMobiles, "at", "", "以逗号分隔的手机号列表用于@")
	flag.BoolVar(&atAll, "atall", false, "@所有人")
	flag.Parse()

	if accessToken == "" || text == "" {
		log.Fatal("token 和 text 为必填")
	}

	client := dingtalk.NewClient(accessToken, secret)

	mobiles := []string{}
	if atMobiles != "" {
		mobiles = strings.Split(atMobiles, ",")
	}

	switch msgType {
	case "markdown":
		msg := dingtalk.NewMarkdownMessage().SetMarkdown(title, text).SetAt(mobiles, atAll)
		if _, err := client.Send(msg); err != nil {
			log.Fatal(err)
		}
	case "text":
		msg := dingtalk.NewTextMessage().SetText(text).SetAt(mobiles, atAll)
		if _, err := client.Send(msg); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("未知类型: %s\n", msgType)
		flag.Usage()
	}
}
