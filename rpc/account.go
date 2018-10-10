package rpc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ripple/crypto"
	"ripple/tools/http"
)

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

func (c *Client) GenAddress() (*crypto.EcdsaKey, string, error) {
	key, err := crypto.GenEcdsaKey()
	if err != nil {
		return key, "", err
	}
	var seq0 uint32
	address, err := crypto.AccountId(key, &seq0)
	return key, address.String(), err
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

type resAccountInfo struct {
	Result resResult
}

type resResult struct {
	Validated          bool
	Status             string
	LedgerCurrentIndex int64          `json:"ledger_current_index"`
	AccountData        resAccountData `json:"account_data"`
}

type resAccountData struct {
	Index    string
	Sequence int64
}

func (c *Client) GetAccountInfo(address string) (*resAccountInfo, error) {
	params := map[string]interface{}{
		"method": "account_info",
		"params": []map[string]string{
			{
				"account": address,
				"ledgder": "validated",
			},
		},
	}
	str, _ := json.Marshal(params)
	resp, err := http.HttpPost(c.rpcURL, str)
	if err != nil {
		return nil, err
	}
	res := &resAccountInfo{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
