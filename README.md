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

**feishu (library)**
```
package main

import "github.com/as7446/message-notice/pkg/feishu"

func main() {
	client := feishu.NewClient("https://open.feishu.cn/open-apis/bot/v2/hook/xxx", "secret") // secret 可为空

	// 文本 + @user_id
	txt := feishu.NewTextMessage().
		SetText("发布完成").
		MentionUserIDs([]string{"ou_xxx"})
	client.Send(txt)

	// post 富文本
	post := feishu.NewPostMessage().
		SetTitle("发布结果").
		SetTextLines([]string{"服务: api", "状态: 通过"})
	client.Send(post)
}
```

**wecom (library)**
```
package main

import "github.com/as7446/message-notice/pkg/wecom"

func main() {
	client := wecom.NewClient("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx")

	// 文本 + @userid
	text := wecom.NewTextMessage().
		SetText("上线完成").
		SetMentionUserIDs([]string{"zhangsan"}).
		AppendVisibleMentions([]string{"zhangsan"})
	client.Send(text)

	// Markdown
	md := wecom.NewMarkdownMessage().
		SetContent("**服务**: api\n**状态**: 通过").
		AppendMentions([]string{"zhangsan"})
	client.Send(md)
}
```

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

**Mattermost（高级）**
- `@提醒`：
```
go run ./cmd/mattermost \
  -webhook "$WEBHOOK" \
  -text "发布完成" \
  -mentions "alice,bob"   # 会自动在文本前插入 "@alice @bob" 可见提醒
```
- 复杂附件（JSON 字符串或文件）：
```
# JSON 字符串（数组或单对象均可）
go run ./cmd/mattermost \
  -webhook "$WEBHOOK" \
  -text "发布结果" \
  -attachments-json '[{"title":"Build #42","text":"OK","color":"#2ecc71","fields":[{"title":"Service","value":"api","short":true}]}]'

# JSON 文件
cat > attach.json <<'JSON'
[
  {
    "title": "Build #42",
    "pretext": "CI",
    "text": "OK",
    "color": "#2ecc71",
    "fields": [
      {"title": "Service", "value": "api", "short": true},
      {"title": "Region", "value": "us-east-1", "short": true}
    ]
  }
]
JSON

go run ./cmd/mattermost -webhook "$WEBHOOK" -text "发布结果" -attachments-file attach.json
```

**Feishu**
```
go run ./cmd/feishu \
  -webhook "https://open.feishu.cn/open-apis/bot/v2/hook/xxx" \
  -secret "$SECRET" \
  -type text \
  -text "发布完成" \
  -at "ou_xxx,ou_yyy"
```
- 富文本（post）：`go run ./cmd/feishu -webhook "..." -type post -title "发布结果" -text "服务: api\n状态: 通过"`

**WeCom（企业微信）**
```
go run ./cmd/wecom \
  -webhook "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx" \
  -type markdown \
  -text "**发布状态**: 成功" \
  -at-userids "zhangsan,lisi" \
  -at-mobiles "1390000"
```

### 构建与交付
- 使用 Make：`make build-all`（在 `bin/` 下生成四个 CLI）
