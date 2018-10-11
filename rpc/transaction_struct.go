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
	TxBlob              string       `json:"tx_blob"`
	TxJson              SubmitTxJson `json:"tx_json"`
}

type SubmitTxJson struct {
	Account            string
	Amount             string
	Destination        string
	Fee                string
	LastLedgerSequence int64
	Sequence           int64
	SigningPubKey      string
	TransactionType    string
	TxnSignature       string
	Hash               string `json:"hash"`
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
	Validated          bool   `json:"validated"`
	Status             string `json:"status"`
	Error              string
	ErrorMessage       string `json:"error_message"`
	ErrorCode          int64  `json:"error_code"`
	Requeset           *TxRequest
}

type TxRequest struct {
	Command     string
	Transaction string
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

type PaymentResp struct {
	Result   string
	Count    int64
	Marker   string
	Payments []Payment
}

type Payment struct {
	Amount                    string
	DeliveredAomunt           string                   `json:"delivered_amount"`
	DestinationBalanceChanges []*PaymentBalanceChanges `json:"destination_balance_changes"`
	SourceBalanceChanges      []*PaymentBalanceChanges `json:"source_balance_changes"`
	TransacationCost          string                   `json:"transacation_cost"`
	TxIndex                   int64                    `json:"tx_index"`
	Currency                  string
	Destination               string
	ExecutedTime              string `json:"executed_time"`
	Issuer                    string
	LedgerIndex               int64 `json:"ledger_index"`
	Source                    string
	SourceCurrency            string
	TxHash                    string `json:"tx_hash"`
}

type PaymentBalanceChanges struct {
	Counterparty string
	Currency     string
	Value        string
}
