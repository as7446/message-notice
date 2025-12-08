package notice

import (
	"fmt"

	"github.com/as7446/message-notice/pkg"
	"github.com/as7446/message-notice/pkg/dingtalk"
	"github.com/as7446/message-notice/pkg/feishu"
	"github.com/as7446/message-notice/pkg/mattermost"
	"github.com/as7446/message-notice/pkg/wecom"
)

// Channel 表示通用的渠道标识
type Channel string

const (
	ChannelDingTalk   Channel = "dingtalk"
	ChannelMattermost Channel = "mattermost"
	ChannelFeishu     Channel = "feishu"
	ChannelWeCom      Channel = "wecom"
)

// Sender 为统一发送接口，兼容各渠道客户端
type Sender interface {
	Send(message pkg.Message) (*pkg.Response, error)
}

// Config 通用配置，按需填充
// - DingTalk 需提供 Token，Secret 可选
// - Mattermost 需提供 Webhook
// - Feishu 需提供 Webhook，Secret 可选
// - WeCom 需提供 Webhook
type Config struct {
	Channel Channel
	Token   string // 钉钉 token
	Secret  string // 钉钉/飞书签名密钥, 可选
	Webhook string // Mattermost/飞书/企业微信 webhook
}

// NewSender 根据渠道构造对应客户端，便于在上层统一注入
func NewSender(cfg Config) (Sender, error) {
	switch cfg.Channel {
	case ChannelDingTalk:
		if cfg.Token == "" {
			return nil, fmt.Errorf("dingtalk token 不能为空")
		}
		return dingtalk.NewClient(cfg.Token, cfg.Secret), nil
	case ChannelMattermost:
		if cfg.Webhook == "" {
			return nil, fmt.Errorf("mattermost webhook 不能为空")
		}
		return mattermost.NewClient(cfg.Webhook), nil
	case ChannelFeishu:
		if cfg.Webhook == "" {
			return nil, fmt.Errorf("feishu webhook 不能为空")
		}
		return feishu.NewClient(cfg.Webhook, cfg.Secret), nil
	case ChannelWeCom:
		if cfg.Webhook == "" {
			return nil, fmt.Errorf("wecom webhook 不能为空")
		}
		return wecom.NewClient(cfg.Webhook), nil
	default:
		return nil, fmt.Errorf("未知渠道: %s", cfg.Channel)
	}
}
