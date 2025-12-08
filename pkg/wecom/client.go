package wecom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/as7446/message-notice/pkg"
)

// Client 企业微信机器人客户端
type Client struct {
	Webhook    string
	HTTPClient *http.Client
}

func NewClient(webhook string) *Client {
	return &Client{
		Webhook:    webhook,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) Send(message pkg.Message) (*pkg.Response, error) {
	res := &pkg.Response{}
	if c.Webhook == "" {
		return res, errors.New("webhook 不能为空")
	}

	b, err := message.ToBytes()
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Webhook, bytes.NewReader(b))
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := c.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	if len(bodyBytes) == 0 && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return res, nil
	}

	if err := json.Unmarshal(bodyBytes, res); err == nil {
		if res.ErrCode != 0 {
			return res, fmt.Errorf("wecom error: code=%d msg=%s", res.ErrCode, res.ErrMsg)
		}
		return res, nil
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return res, nil
	}

	return res, errors.New("发送企业微信消息失败")
}
