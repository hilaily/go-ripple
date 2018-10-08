package rpc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ripple/tools/http"
)

type Client struct {
	rpcURL string
	apiURL string
}

func NewClient(rpcURL, apiURL string) *Client {
	return &Client{
		rpcURL: rpcURL,
		apiURL: apiURL,
	}
}

type AccountBalancesStruct struct {
	Result      string     `json: "result"`
	LedgerIndex int64      `json: "ledger_index"`
	CloseTime   string     `json: "close_time"`
	Limit       int        `json: "limit"`
	Balances    []*Balance `json: "balances"`
	Message     string
}

type Balance struct {
	Currency     string `json:"currency"`
	Counterparty string
	Value        string
}

// GetAccountBalances 获取账户余额
func (c *Client) GetAccountBalances(address string, queryParams map[string]string) (*AccountBalancesStruct, error) {
	balance := &AccountBalancesStruct{}
	if address == "" {
		return balance, fmt.Errorf("address is empty")
	}
	host := "/v2/accounts/" + address + "/balances"
	values := make(url.Values)
	for key, val := range queryParams {
		values.Add(key, val)
	}
	queryUrl := c.apiURL + host + values.Encode()
	resp, err := http.HttpGet(queryUrl)
	if err != nil {
		return balance, err
	}
	err = json.Unmarshal(resp, balance)
	if err != nil {
		return balance, err
	}
	if balance.Result != "success" {
		return balance, fmt.Errorf(balance.Message)
	}
	return balance, nil
}
