package rpc

import (
	"encoding/hex"
	"fmt"
	"ripple/crypto"
	"ripple/data"
)

const (
	TXN_TYPE_PAYMENT = "Payment"
)

type Transaction struct {
	TransactionType string
	Account         string
	Destination     string
	Amount          *Amount
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
	txn := &Transaction{
		TransactionType: "Payment",
		Account:         from,
		Destination:     to,
		Amount: &Amount{
			Currency: currency,
			Value:    value,
			Issuer:   from,
		},
	}
	c.SignOffline(txn, privateKey)

	params := Params{
		Method: "sign",
		Params: []map[string]interface{}{
			"offline":      true,
			"tx_json":      txn,
			"fee_mult_max": 1000,
		},
	}
	return nil
}

// Sign 给交易签名
// 使用的这个库的方法，没有完全理解其逻辑
func (c *Client) SignOffline(txn *Transaction, privateKey string) error {
	fromAccount, _ := data.NewAccountFromAddress(txn.Account)
	toAccount, _ := data.NewAccountFromAddress(txn.Destination)
	amount, _ := data.NewAmount(txn.Amount.Value + "/" + txn.Amount.Currency)
	txnBase := data.TxBase{
		TransactionType: data.PAYMENT,
		Account:         *fromAccount,
	}
	payment := data.Payment{
		TxBase:      txnBase,
		Destination: *toAccount,
		Amount:      *amount,
	}

	pri, _ := hex.DecodeString(privateKey)
	key := crypto.LoadECDSKey(pri)

	fmt.Println("pri: ", hex.EncodeToString(key.D.Bytes()))
	seq := uint32(1)
	err := data.Sign(&payment, key, &seq)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", payment)
	return nil
}

func (c *Client) Submit() {}
