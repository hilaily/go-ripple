package rpc

import (
	"fmt"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	res, err := client.GetServerInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("%+v", res))
}
