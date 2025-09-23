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

	flag.StringVar(&accessToken, "token", "", "DingTalk robot access token")
	flag.StringVar(&secret, "secret", "", "DingTalk robot secret (optional)")
	flag.StringVar(&msgType, "type", "text", "message type: text|markdown")
	flag.StringVar(&title, "title", "", "markdown title (for markdown type)")
	flag.StringVar(&text, "text", "", "message content")
	flag.StringVar(&atMobiles, "at", "", "comma-separated mobile numbers to @")
	flag.BoolVar(&atAll, "atall", false, "@ all members")
	flag.Parse()

	if accessToken == "" || text == "" {
		log.Fatal("token and text are required")
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
		fmt.Printf("unknown type: %s\n", msgType)
		flag.Usage()
	}
}
