package rpc

import (
	"encoding/hex"
	"fmt"
	"ripple/crypto"
	"testing"
)

func TestSign(t *testing.T) {
	from := "rsHRq2asWJDs87eRgXNUCZk2PjbfrPgUKh"
	to := "rh4WZwXaDhamjM7hw8gArB9Jgs6fkxUGnw"
	value := "0.1"
	currency := "XRP"
	key, _ := crypto.NewECDSAKey([]byte("1"))
	k := hex.EncodeToString(key.D.Bytes())
	err := client.Transfer(from, to, currency, value, k)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
