/*
@Author : Fuxuhao
@Time : 2024/7/26 18:52
*/
package mattermost

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

type Client struct {
	webhookUrl string
	HTTPClient *http.Client
}

func NewClient(webhook string) *Client {
	return &Client{webhookUrl: webhook, HTTPClient: &http.Client{Timeout: 10 * time.Second}}
}
func (c *Client) Send(message pkg.Message) (*pkg.Response, error) {
	res := &pkg.Response{}
	b, err := message.ToBytes()
	if err != nil {
		return res, err
	}
	req, err := http.NewRequest(http.MethodPost, c.webhookUrl, bytes.NewReader(b))
	if err != nil {
		return res, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := c.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	// Mattermost 入站 webhook 通常返回 200 OK 且响应体为空
	if len(respBytes) == 0 && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return res, nil
	}
	if err = json.Unmarshal(respBytes, &res); err == nil {
		if res.ErrCode != 0 {
			return res, fmt.Errorf("mattermost error: code=%d msg=%s", res.ErrCode, res.ErrMsg)
		}
		return res, nil
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return res, nil
	}
	return res, errors.New("send message mattermost failed")
}
