package rpc

import (
	"fmt"
	"ripple/tools/http"
)

func (c *Client) GetServerInfo() error {
	resp, err := http.HttpPost(c.rpcURL, []byte(`{"method":"server_state", "params": [{}]}`))
	if err != nil {
		return err
	}
	fmt.Println("resp: ", string(resp))
	return nil
}
