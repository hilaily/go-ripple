package rpc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ripple/tools/http"
	"strconv"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
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

func (c *Client) GetAccountBalances(address, currency string, date *time.Time, limit int) (*AccountBalancesStruct, error) {
	balance := &AccountBalancesStruct{}
	if address == "" {
		return balance, fmt.Errorf("address is empty")
	}
	host := "/v2/accounts/" + address + "/balances"
	values := make(url.Values)
	if currency != "" {
		values.Add("currency", currency)
	}
	if date != nil {
		values.Add("date", date.Format("2006-01-02T15:04:05Z"))
	}
	if limit > 0 {
		values.Add("limit", strconv.Itoa(3))
	}
	queryUrl := c.url + host + values.Encode()
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
