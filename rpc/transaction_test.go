package rpc

import (
	"encoding/hex"
	"fmt"
	"go-ripple/data"
	"testing"
)

func TestSign(t *testing.T) {
	//from := "rsHRq2asWJDs87eRgXNUCZk2PjbfrPgUKh"
	//to := "rBx7ozXTBtEzXw6FajXAAyGKG5G5b8fhh3"
	//value := "0.1"
	//currency := "XRP"
	//pri := "86029426A6D950A14CEDD1AE33F0EB8C7CE1C0E8190D41D82C52EA160084B9E8"

	//_, err := client.Transfer(from, to, currency, value, pri)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//}
}

// tx_blob
// 1200002280000000240000000F201B00CB247E6140000000000186A068400000000000000C7321028C35EEA94EE7FA9C8485426E164159330BA2453368F399669D5110009F270EE974473045022100D59891D15129AFA2297506207AF14A97C2C236C690BA5E167E84BC070CA3774202203F80DFC3D8965AA4705940B9233ED8570812557F1E9DC011DEAF47DC2AE8BD588114190BA3E39BE7E7267AF0C79B2E3E2BDEC738A154831424F9C8900B8C33E55A2B848587884B10EF9992C7
func TestMakeBlob(t *testing.T) {
	Account := "rsHRq2asWJDs87eRgXNUCZk2PjbfrPgUKh"
	Amount := "0.1"
	Destination := "rh4WZwXaDhamjM7hw8gArB9Jgs6fkxUGnw"
	Fee := "12"
	Flags := 2147483648
	last := uint32(13313150)
	Sequence := 15
	TxnSignature := "3045022100D59891D15129AFA2297506207AF14A97C2C236C690BA5E167E84BC070CA3774202203F80DFC3D8965AA4705940B9233ED8570812557F1E9DC011DEAF47DC2AE8BD58"
	SigningPubKey :=
		"028C35EEA94EE7FA9C8485426E164159330BA2453368F399669D5110009F270EE9"

	fromAccount, _ := data.NewAccountFromAddress(Account)
	toAccount, _ := data.NewAccountFromAddress(Destination)
	amount, _ := data.NewAmount(Amount + "/XRP")
	fee, _ := data.NewValue(Fee, true)
	flags := data.TransactionFlag(Flags)
	tSig, _ := hex.DecodeString(TxnSignature)
	txnSign := data.VariableLength(tSig)
	signPubKey := data.PublicKey{}
	pk, _ := hex.DecodeString(SigningPubKey)
	copy(signPubKey[:], pk)

	txn := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		LastLedgerSequence: &last,
		Flags:              &flags,
		Sequence:           uint32(Sequence),
		TxnSignature:       &txnSign,
		Fee:                *fee,
		SigningPubKey:      &signPubKey,
	}
	payment := data.Payment{
		TxBase:      txn,
		Amount:      *amount,
		Destination: *toAccount,
	}

	res, err := client.makeTxBlob(&payment)
	if err != nil {
		t.Error("gen blob err: ", err)
	}
	t.Log("tx blog: ", res)
}

func TestTX(t *testing.T) {
	hash := "A2603189C714F39CBBDD29705360EA3EA8EDECAD5FF6FE2762E65814A9408151"
	res, err := client.TX(hash)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
