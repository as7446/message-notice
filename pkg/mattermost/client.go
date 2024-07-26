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
	"github.com/as7446/message-notice/pkg"
	"io"
	"net/http"
)

type Client struct {
	webhookUrl string
}

func NewClient(webhook string) *Client {
	return &Client{webhookUrl: webhook}
}
func (c *Client) Send(message pkg.Message) (*pkg.Response, error) {
	res := &pkg.Response{}
	b, err := message.ToBytes()
	if err != nil {
		return res, err
	}
	req, err := http.NewRequest(http.MethodPost, c.webhookUrl, bytes.NewReader(b))
	if err != nil {
		return res, nil
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	fmt.Println(string(respBytes))
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return res, err
	}
	if res.ErrCode != 0 {
		return res, errors.New("send message mattermost failed")
	}
	return res, nil
}
