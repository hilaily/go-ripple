package rpc

import (
	"fmt"
	"testing"
)

var (
	//client = NewClient("https://s.altnet.rippletest.net:51234/v2")
	client = NewClient("https://data.ripple.com")
)

func TestGetAccountBalance(t *testing.T) {
	address := "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"
	res, err := client.GetAccountBalances(address, "", nil, 0)
	if err != nil {
		t.Error("get err: ", err)
	}
	for _, v := range res.Balances {
		fmt.Printf("balance: %+v\n", v)
	}
}
