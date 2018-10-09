package rpc

import "testing"

func TestGetServerInfo(t *testing.T) {
	err := client.GetServerInfo()
	if err != nil {
		t.Error(err)
	}

}
