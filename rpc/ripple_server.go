package rpc

import (
	"encoding/json"
	"ripple/tools/http"
)

func (c *Client) GetServerInfo() (*ServerInfoResult, error) {
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(`{"method":"server_state", "params": [{}]}`))
	if err != nil {
		return nil, err
	}
	res := &ServerInfoResp{}
	err = json.Unmarshal(resp, res)
	return res.Result, err
}
