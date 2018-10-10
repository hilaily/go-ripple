package rpc

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
