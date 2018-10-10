package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-ripple/crypto"
	"go-ripple/data"
	"go-ripple/tools/http"
)

const (
	default_currency = "XRP"
)

// sign 命令
// https://developers.ripple.com/sign.html
// Sign 给交易签名
// 使用的这个库的方法，没有完全理解其逻辑
//
// Transfer 发起交易
// from, to 账户地址
// currency 货币类型 默认 XRP
// value 交易金额
// privateKey 私钥的 16 进制编码
// accountSequence 账户的 sequence
// lastLedgerSequence 交易允许的最大账本 https://developers.ripple.com/reliable-transaction-submission.html#lastledgersequence
func (c *Client) Sign(from, to, currency, value, fee, privateKey string, accountSequence, lastLedgerSequence uint32) (string, error) {
	fromAccount, _ := data.NewAccountFromAddress(from)
	toAccount, _ := data.NewAccountFromAddress(to)
	amount, _ := data.NewAmount(value + "/" + currency)
	feeVal, _ := data.NewValue(fee, true)

	txnBase := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		Sequence:           accountSequence,
		Fee:                *feeVal,
		LastLedgerSequence: &lastLedgerSequence,
	}
	payment := &data.Payment{
		TxBase:      txnBase,
		Destination: *toAccount,
		Amount:      *amount,
	}

	txBlob, err := c.signOffline(payment, privateKey)
	if err != nil {
		return "", err
	}
	return txBlob, nil
}

func (c *Client) signOffline(payment *data.Payment, privateKey string) (string, error) {
	pri, _ := hex.DecodeString(privateKey)
	key := crypto.LoadECDSKey(pri)

	err := data.Sign(payment, key, nil)
	if err != nil {
		return "", err
	}
	return c.makeTxBlob(payment)
}

// MakeTxblob
// 构造 txBlob，用于之后提交交易
func (c *Client) makeTxBlob(payment *data.Payment) (string, error) {
	fmt.Println("sign pub key: ", payment.SigningPubKey.String())
	_, raw, err := data.Raw(data.Transaction(payment))
	if err != nil {
		return "", err
	}
	txBlob := fmt.Sprintf("%X", raw)
	return txBlob, nil
}

// submit 命令
// https://developers.ripple.com/submit.html
// Submit ripple submit command
// 提交交易给瑞波链
func (c *Client) Submit(txBlob string) (*SubmitResult, error) {
	res := &SubmitResp{}
	params := `{"method": "submit", "params": [{"tx_blob": "` + txBlob + `"}]}`
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(params))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, res)
	return res.Result, err
}

// tx 命令
// https://developers.ripple.com/tx.html
func (c *Client) TX(hash string) (*TxResult, error) {
	params := `{"method":"tx", "params": [{"transaction":"` + hash + `"}]}`
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(params))
	if err != nil {
		return nil, err
	}
	res := &TxResp{}
	err = json.Unmarshal(resp, res)
	return res.Result, nil
}
