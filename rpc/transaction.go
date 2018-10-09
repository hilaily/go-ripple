package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"ripple/crypto"
	"ripple/data"
	"ripple/tools/http"
)

const (
	TXN_TYPE_PAYMENT = "Payment"
)

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

func (c *Client) Transfer(from, to, currency, value, privateKey string) error {
	accountInfo, err := c.GetAccountInfo(from)
	if err != nil {
		return fmt.Errorf("get account info err: %v\n", err)
	}

	fee := "12"
	seq := uint32(accountInfo.Result.AccountData.Sequence + 1)
	lastLedgerSequence := uint32(accountInfo.Result.LedgerCurrentIndex + 5)

	fromAccount, _ := data.NewAccountFromAddress(from)
	toAccount, _ := data.NewAccountFromAddress(to)
	amount, _ := data.NewAmount(value + "/" + currency)
	feeVal, _ := data.NewValue(fee, true)

	//flags := data.TransactionFlag(uint32(2147483648))
	//seq = uint32(10)
	//lastLedgerSequence = uint32(13308150)

	txnBase := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		Sequence:           seq,
		Fee:                *feeVal,
		LastLedgerSequence: &lastLedgerSequence,
	}
	payment := &data.Payment{
		TxBase:      txnBase,
		Destination: *toAccount,
		Amount:      *amount,
	}

	txBlob, err := c.SignOffline(payment, privateKey)
	if err != nil {
		return err
	}

	py, _ := json.Marshal(payment)
	fmt.Printf("resp: %s\n", py)
	fmt.Println("tx blob: ", txBlob)

	// submit a transaction
	params := `{"method": "submit", "params": [{"tx_blob": "` + txBlob + `"}]}`
	resp, err := http.HttpPost(c.rpcURL, []byte(params))
	if err != nil {
		return err
	}
	fmt.Printf("resp: %s\n", string(resp))
	return nil
}

// Sign 给交易签名
// 使用的这个库的方法，没有完全理解其逻辑
func (c *Client) SignOffline(payment *data.Payment, privateKey string) (string, error) {
	pri, _ := hex.DecodeString(privateKey)
	key := crypto.LoadECDSKey(pri)

	err := data.Sign(payment, key, &payment.Sequence)
	if err != nil {
		return "", err
	}
	return c.MakeTxBlob(payment)
}

func (c *Client) MakeTxBlob(payment *data.Payment) (string, error) {
	fmt.Println("sign pub key: ", payment.SigningPubKey.String())
	_, raw, err := data.Raw(data.Transaction(payment))
	if err != nil {
		return "", err
	}
	txBlob := fmt.Sprintf("%X", raw)
	return txBlob, nil
}

func (c *Client) Submit() {

}
