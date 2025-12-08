package feishu

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/as7446/message-notice/pkg"
)

// Client 飞书机器人客户端
type Client struct {
	Webhook      string
	Secret       string
	HTTPClient   *http.Client
	NowTimestamp func() int64
}

// NewClient 创建客户端；secret 可为空（无需签名）
func NewClient(webhook, secret string) *Client {
	return &Client{
		Webhook:      webhook,
		Secret:       secret,
		HTTPClient:   &http.Client{Timeout: 10 * time.Second},
		NowTimestamp: func() int64 { return time.Now().Unix() },
	}
}

// Send 发送消息；如配置了 secret，会自动生成签名并注入消息体
func (c *Client) Send(message pkg.Message) (*pkg.Response, error) {
	res := &pkg.Response{}
	if c.Webhook == "" {
		return res, errors.New("webhook 不能为空")
	}

	if c.Secret != "" {
		signable, ok := message.(signer)
		if !ok {
			return res, errors.New("当前消息类型不支持签名，请使用 feishu.TextMessage 或 feishu.PostMessage")
		}
		ts := c.NowTimestamp()
		sign, err := c.generateSign(ts)
		if err != nil {
			return res, err
		}
		signable.applySignature(ts, sign)
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

	// 优先解析通用响应
	if err := json.Unmarshal(bodyBytes, res); err == nil {
		if res.ErrCode != 0 {
			return res, fmt.Errorf("feishu error: code=%d msg=%s", res.ErrCode, res.ErrMsg)
		}
		return res, nil
	}

	// 兼容飞书返回字段
	var fr feishuResp
	if err := json.Unmarshal(bodyBytes, &fr); err == nil {
		mapped := fr.toResponse()
		if mapped.ErrCode != 0 {
			return mapped, fmt.Errorf("feishu error: code=%d msg=%s", mapped.ErrCode, mapped.ErrMsg)
		}
		return mapped, nil
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return res, nil
	}
	return res, errors.New("发送飞书消息失败")
}

// generateSign 生成签名
func (c *Client) generateSign(ts int64) (string, error) {
	if c.Secret == "" {
		return "", nil
	}
	stringToSign := fmt.Sprintf("%d\n%s", ts, c.Secret)
	h := hmac.New(sha256.New, []byte(c.Secret))
	if _, err := h.Write([]byte(stringToSign)); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

type feishuResp struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
	Code          int    `json:"code"`
	Msg           string `json:"msg"`
}

func (r feishuResp) toResponse() *pkg.Response {
	code := r.Code
	if code == 0 {
		code = r.StatusCode
	}
	msg := r.Msg
	if msg == "" {
		msg = r.StatusMessage
	}
	return &pkg.Response{ErrCode: code, ErrMsg: msg}
}
