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
	var attach string // simple format: title|text|color
	var attachmentsJSON string
	var attachmentsFile string
	var mentions string // comma-separated usernames to @ in text

	flag.StringVar(&webhook, "webhook", "", "Mattermost incoming webhook URL")
	flag.StringVar(&text, "text", "", "message content")
	flag.StringVar(&channel, "channel", "", "target channel name or id")
	flag.StringVar(&username, "username", "", "override username")
	flag.StringVar(&iconURL, "icon-url", "", "override icon url")
	flag.StringVar(&iconEmoji, "icon-emoji", "", "override icon emoji, e.g. :rocket:")
	flag.StringVar(&attach, "attach", "", "attachment in 'title|text|color' format; repeatable via multiple runs")
	flag.StringVar(&attachmentsJSON, "attachments-json", "", "JSON string for full-featured attachments (array or single object)")
	flag.StringVar(&attachmentsFile, "attachments-file", "", "Path to JSON file containing attachments (array or single object)")
	flag.StringVar(&mentions, "mentions", "", "comma-separated usernames to @ in text (e.g. user1,user2)")
	flag.Parse()

	if webhook == "" || text == "" {
		log.Fatal("webhook and text are required")
	}

	// prepend mentions into text if provided
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

	// parse attachments from JSON string
	if attachmentsJSON != "" {
		addAttachmentsFromJSON(msg, attachmentsJSON)
	}
	// parse attachments from JSON file
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

// addAttachmentsFromJSON supports either a single object or an array of attachments
func addAttachmentsFromJSON(msg *mattermost.TextMessage, jsonStr string) {
	jsonStr = strings.TrimSpace(jsonStr)
	if jsonStr == "" {
		return
	}
	// Try array first
	var arr []mattermost.Attachment
	if err := json.Unmarshal([]byte(jsonStr), &arr); err == nil {
		for _, a := range arr {
			msg.AddAttachment(a)
		}
		return
	}
	// Try single object
	var one mattermost.Attachment
	if err := json.Unmarshal([]byte(jsonStr), &one); err == nil {
		msg.AddAttachment(one)
	}
}
