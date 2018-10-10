package rpc

type AccountInfoResp struct {
	Result *AccountInfoResult
}

type AccountInfoResult struct {
	Validated          bool
	Status             string
	LedgerCurrentIndex int64            `json:"ledger_current_index"`
	AccountData        *AccountInfoData `json:"account_data"`
}

type AccountInfoData struct {
	Index    string
	Sequence int64
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
