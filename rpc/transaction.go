package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"ripple/crypto"
	"ripple/data"
	"ripple/tools/http"
	"strconv"
)

const (
	default_currency = "XRP"
)

// Transfer 发起交易
// from, to 账户地址
// currency 货币类型 默认 XRP
// value 交易金额
// privateKey 私钥的 16 进制编码
func (c *Client) Transfer(from, to, currency, value, privateKey string) (*SubmitResult, error) {
	// 获取账户的 Sequence 和 LedgerCurrentIndex
	// 交易流程 https://developers.ripple.com/reliable-transaction-submission.html
	accountInfo, err := c.GetAccountInfo(from)
	if err != nil {
		return nil, err
	}
	if currency == "" {
		currency = default_currency
	}

	serverInfo, err := c.GetServerInfo()
	if err != nil {
		return nil, err
	}

	fee := strconv.FormatInt(serverInfo.State.ValidatedLedger.BaseFee, 10)
	seq := uint32(accountInfo.Result.AccountData.Sequence)
	lastLedgerSequence := uint32(accountInfo.Result.LedgerCurrentIndex + 5)

	fromAccount, _ := data.NewAccountFromAddress(from)
	toAccount, _ := data.NewAccountFromAddress(to)
	amount, _ := data.NewAmount(value + "/" + currency)
	feeVal, _ := data.NewValue(fee, true)

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
		return nil, err
	}
	resp, err := c.Submit(txBlob)
	if err != nil {
		return nil, err
	}
	if resp.EngineResultCode != 0 {
		return resp, fmt.Errorf(resp.EngineResultMessage)
	}
	return resp, nil
}

// sign 命令
// https://developers.ripple.com/sign.html
// Sign 给交易签名
// 使用的这个库的方法，没有完全理解其逻辑
func (c *Client) SignOffline(payment *data.Payment, privateKey string) (string, error) {
	pri, _ := hex.DecodeString(privateKey)
	key := crypto.LoadECDSKey(pri)

	err := data.Sign(payment, key, nil)
	if err != nil {
		return "", err
	}
	return c.MakeTxBlob(payment)
}

// MakeTxblob
// 构造 txBlob，用于之后提交交易
func (c *Client) MakeTxBlob(payment *data.Payment) (string, error) {
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
