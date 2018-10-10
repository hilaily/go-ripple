package rpc

type SubmitResp struct {
	Result *SubmitResult
}

// submit 命令的返回
type SubmitResult struct {
	EngineResult        string `json:"engine_result"`
	EngineResultCode    int64  `json:"engine_result_code"`
	EngineResultMessage string `json:"engine_result_message"`
	Status              string
	TxBlob              string `json:"tx_blob"`
	TxJson              string `json:"tx_json"`
	Destination         string
	Fee                 string
	Flags               int64
	Sequence            int64
	SigningPubKey       string
	TransactionTyep     string
	TxnSignature        string
	Hash                string
}

type TxJson struct {
	Account string
	Amount  Amount
}

type Transaction struct {
	TransactionType    string
	Account            string
	Destination        string
	Amount             *Amount
	Sequence           int64
	LastLedgerSequence int64
	Fee                int64
}

type Amount struct {
	Currency string
	Value    string
	Issuer   string
}

type Params struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// tx 命令的返回
type TxResp struct {
	Result *TxResult
}

type TxResult struct {
	Account            string
	Amount             string
	Destination        string
	Fee                string
	LastLedgerSequence int64
	Sequence           int64
	SigningPubKey      string
	TransactionType    string
	TxnSignatrue       string
	Date               int64
	Hash               string
	InLedger           int64
	LedgerIndex        int64 `json:"ledger_index"`
	Meta               *TxMeta
	Status             string `json:"status"`
	Validated          bool   `json:"validated"`
}

type TxMeta struct {
	AffectedNodes     []TxModifyNode
	TransactionIndex  int64
	TransactionResult string
	DeliveredAomunt   string `json:"delivered_amount"`
}

type TxModifyNode struct {
	FinalFields     TxFinalFields
	LedgerEntryType string
	LedgerIndex     string
	PreviousFields  TxPreviousFields
}

type TxFinalFields struct {
	Account    string
	Balance    string
	Flags      int64
	OwnerCount int64
	Sequence   int64
}

type TxPreviousFields struct {
	Balance  string
	Sequence int64
}
