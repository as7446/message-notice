## message-notice
### 快速开始
**安装**
```
go get github.com/as7446/message-notice/pkg
```
**dingtalk (library)**
```
package main

import (
	"github.com/as7446/message-notice/pkg/dingtalk"
)
func main() {
	dingClient := dingtalk.NewClient("accessToken", "secret")
	// 文本消息 + 按手机号@
	txt := dingtalk.NewTextMessage().SetText("测试@提醒 @1510000").SetAt([]string{"1510000"}, false)
	dingClient.Send(txt)
	// Markdown 消息 + 按用户ID@
	md := dingtalk.NewMarkdownMessage().SetMarkdown("标题", "内容中可写 @1510000 便于可视化").SetAtUserIds([]string{"abcd1234"}, false)
	dingClient.Send(md)
}
```
说明：要在会话中显示可见的@，请在文本/Markdown内容里手动写入相应的 `@手机号`；`at` 字段用于触发提醒。

**mattermost (library)**
```
package main

import (
	"github.com/as7446/message-notice/pkg/mattermost"
)

func main() {
	mattermostClient := mattermost.NewClient("https://xxxxxx")
	msg := mattermost.NewTextMessage().
		SetText("Hello from webhook").
		SetChannel("town-square").
		SetUsername("bot").
		SetIconEmoji(":rocket:").
		AddAttachment(mattermost.Attachment{Title: "Build", Text: "OK", Color: "#2ecc71"})
	mattermostClient.Send(msg)
}
```
- 支持字段：`text`、`channel`、`username`、`icon_url`、`icon_emoji`、`attachments`、`props`
- 附件字段：`fallback`、`color`、`pretext`、`author_*`、`title`、`title_link`、`text`、`fields`、`image_url`、`thumb_url`、`footer`、`footer_icon`、`ts`

### CLI 使用
**DingTalk**
```
go run ./cmd/dingtalk \
  -token "$TOKEN" -secret "$SECRET" \
  -type markdown -title "标题" -text "内容，包含 @1510000 以便可见" \
  -at "1510000,1510001" -atall=false
```
- **文本**: `go run ./cmd/dingtalk -token "$TOKEN" -text "纯文本消息 @1510000" -at "1510000"`
- **用户ID@**: 当前 CLI 仅支持手机号@。如需 `userId` 支持，可在库中使用 `SetAtUserIds`。

**Mattermost**
```
go run ./cmd/mattermost \
  -webhook "$WEBHOOK" \
  -text "Hello from webhook" \
  -channel "town-square" \
  -username "bot" \
  -icon-emoji ":rocket:" \
  -attach "Build|OK|#2ecc71"
```
- `-icon-url` 可用图片 URL 替代 `-icon-emoji`
- `-attach` 简易格式：`title|text|color`（如需复杂附件，请在库侧使用 `AddAttachment` 并填充更多字段）
