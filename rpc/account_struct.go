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

type AccountTransactionResp struct {
	Result       string
	Count        int
	Marker       string
	Transactions []*AccountTransactionData
}

type AccountTransactionData struct {
	Hash        string
	LedgerIndex int64 `json:"ledger_index"`
	Date        string
	Tx          *AccountTransactionTx
	Meta        *AccountTransactionMeta
}

type AccountTransactionTx struct {
	TransactionType    string
	Flags              int64
	Sequence           int64
	LastLedgerSequence int64
	Amount             string
	Fee                string
	SigningPubKey      string
	TxnSignature       string
	Account            string
	Destination        string
	DestinationTag     int64
}

type AccountTransactionMeta struct {
	TransactionIndex int64
	// AffectedNodes     []*AccountTransactionAffectedNodes
	TransactionResult string
	DeliveredAmount   string `json:"delivered_amount"`
}
