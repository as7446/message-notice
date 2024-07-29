## message-notice
钉钉、mattermost使用markdown格式文本发送消息，后续会继续完善钉钉text、link等类型；mattermost目前只能简单text发送。
### 快速开始
安装
```
go get github.com/as7446/message-notice/pkg
```
#### 使用
**dingtalk**
```
package main

import (
	"github.com/as7446/message-notice/pkg/dingtalk"
)
func main() {
	dingClient := dingtalk.NewClient("accessToken", "secret")
	msg := dingtalk.NewMarkdownMessage().SetMarkdown("测试", "测试dingtalk发送 **注意@!**").SetAt([]string{"151xxxx"}, false)
	dingClient.Send(msg)
}
```
**mattermost**
```
package main

import (
	"github.com/as7446/message-notice/pkg/mattermost"
)

func main() {
	mattermostClient := mattermost.NewClient("https://chat.service.shujiajia.com/hooks/sep4fnmedtne5rrjfi8miywzzh")
	msg := mattermost.NewTextMessage().SetText("测试发送 **注意")
	mattermostClient.Send(msg)
}
```
