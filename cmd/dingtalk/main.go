/*
@Author : Fuxuhao
@Time : 2024/7/26 11:03
*/
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/as7446/message-notice/pkg/dingtalk"
	"io"
)

func sign(timestamp int64, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// https://oapi.dingtalk.com/robot/send?access_token=6813a35d4da80b19f146afc2c0e11efdfdbc20cd42318d3de3f3d75a4fd053c7

func main() {
	dingClient := dingtalk.NewClient("6813a35d4da80b19f146afc2c0e11efdfdbc20cd42318d3de3f3d75a4fd053c7", "SEC433b31ffbbd6ebe71c24f1fdfe520bbb27de3c6ef91f6e8791b8f4eb55c1fe74")
	msg := dingtalk.NewMarkdownMessage().SetMarkdown("test", "测试文本&at 某个人").SetAt([]string{"15120027019"}, false)
	fmt.Printf("%#v\n", msg)
	send, err := dingClient.Send(msg)
	fmt.Println(send)
	fmt.Println(err)
}
