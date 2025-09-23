/*
@Author : Fuxuhao
@Time : 2024/7/26 10:51
*/
package dingtalk

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
	"net/url"
	"strconv"
	"time"

	"github.com/as7446/message-notice/pkg"
)

func URL(accessToken, secret string) (string, error) {
	timestamp := time.Now().Unix() * 1000
	return URLWithTimestamp(accessToken, secret, timestamp)
}

func URLWithTimestamp(accessToken, secret string, timestamp int64) (string, error) {
	if accessToken == "" {
		return "", errors.New("access token is empty")
	}
	dingtalkUrl := url.URL{
		Scheme: "https",
		Path:   "robot/send",
		Host:   "oapi.dingtalk.com",
	}
	value := url.Values{}
	value.Set("access_token", accessToken)
	if secret == "" {
		dingtalkUrl.RawQuery = value.Encode()
		return dingtalkUrl.String(), nil
	}
	h := hmac.New(sha256.New, []byte(secret))
	StringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	h.Write([]byte(StringToSign))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	value.Set("timestamp", strconv.FormatInt(timestamp, 10))
	// 交由 url.Values.Encode 处理 URL 编码
	value.Set("sign", sign)
	dingtalkUrl.RawQuery = value.Encode()

	return dingtalkUrl.String(), nil
}

type Client struct {
	AccessToken string
	Secret      string
	HTTPClient  *http.Client
}

func NewClient(accessToken, secret string) *Client {
	return &Client{AccessToken: accessToken, Secret: secret, HTTPClient: &http.Client{Timeout: 10 * time.Second}}
}

func (c *Client) Send(message pkg.Message) (*pkg.Response, error) {
	res := &pkg.Response{}
	b, err := message.ToBytes()
	if err != nil {
		return res, err
	}
	dingUrl, err := URL(c.AccessToken, c.Secret)
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(http.MethodPost, dingUrl, bytes.NewReader(b))
	if err != nil {
		return res, err
	}
	req.Header.Add("Accept-Charset", "utf8")
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
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return res, err
	}
	if res.ErrCode != 0 {
		return res, fmt.Errorf("dingtalk error: code=%d msg=%s", res.ErrCode, res.ErrMsg)
	}
	return res, nil
}
