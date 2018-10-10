package rpc

type ServerInfoResp struct {
	Result *ServerInfoResult
}

type ServerInfoResult struct {
	State  ServerInfoState
	Status string
}

type ServerInfoState struct {
	BuildVersion    string                    `json:"build_version"`
	CompleteLedgers string                    `json:"complete_ledgers"`
	ValidatedLedger ServerInfoValidatedLedger `json:"validated_ledger"`
}

type ServerInfoValidatedLedger struct {
	BaseFee     int64 `json:"base_fee"`
	CloseTime   int64 `json:"close_time"`
	Hash        string
	ReserveBase int64 `json:"reserve_base"`
	ReserveInc  int64 `json:"reserve_inc"`
	Seq         int64
}
