/*
@Author : Fuxuhao
@Time : 2024/7/26 12:27
*/
package pkg

type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

type Notice interface {
	Send() (*Response, error)
}

type Message interface {
	ToBytes() ([]byte, error)
}
