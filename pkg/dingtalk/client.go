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
	"github.com/as7446/message-notice/pkg"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	StringToSign := fmt.Sprintf("%d\n%s", timestamp, accessToken)
	fmt.Println("sign: ", StringToSign)
	if _, err := io.WriteString(h, StringToSign); err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	value.Set("timestamp", strconv.FormatInt(timestamp, 10))
	value.Set("sign", sign)
	dingtalkUrl.RawQuery = value.Encode()
	fmt.Println("dingtalk url: ", dingtalkUrl.String())
	return dingtalkUrl.String(), nil
}

type Client struct {
	AccessToken string
	Secret      string
}

func NewClient(accessToken, secret string) *Client {
	return &Client{accessToken, secret}
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
		return res, nil
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return res, err
	}
	if res.ErrCode != 0 {
		return res, errors.New("send message dingtalk failed")
	}
	return res, nil
}
