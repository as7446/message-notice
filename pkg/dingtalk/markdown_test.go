/*
@Author : Fuxuhao
@Time : 2024/7/26 11:32
*/
package dingtalk

import (
	"testing"
)

func TestURL(t *testing.T) {
	access_token := "6813a35d4da80b19f146afc2c0e11efdfdbc20cd42318d3de3f3d75a4fd053c7"
	secret := "SEC433b31ffbbd6ebe71c24f1fdfe520bbb27de3c6ef91f6e8791b8f4eb55c1fe74"
	url, err := URL(access_token, secret)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(url)
}
