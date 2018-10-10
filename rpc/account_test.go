package rpc

import (
	"fmt"
	"testing"
)

var (
	client = NewClient("https://s.altnet.rippletest.net:51234", "https://testnet.data.api.ripple.com")
	//client = NewClient("http://47.75.70.201:9003", "http://47.75.70.201:9003")
	//client = NewClient("https://data.ripple.com")
)

func TestGetAccountBalance(t *testing.T) {
	address := "rh4WZwXaDhamjM7hw8gArB9Jgs6fkxUGnw"
	res, err := client.GetAccountBalances(address, map[string]string{})
	if err != nil {
		t.Error("get err: ", err)
	}
	for _, v := range res.Balances {
		fmt.Printf("balance: %+v\n", v)
	}
}

func TestGetAccountInfo(t *testing.T) {
	address := "rh4WZwXaDhamjM7hw8gArB9Jgs6fkxUGnw"
	res, err := client.GetAccountInfo(address)
	if err != nil {
		t.Error("err: ", err)
	}
	fmt.Printf("res: %+v\n", res)
}

func TestGenAddress(t *testing.T) {
	key, addr, err := client.GenAddress()
	if err != nil {
		t.Error(err)
	}
	var seq0 uint32

	t.Log("pub: ", key.Public(&seq0))
	t.Log("pri: ", key.Private(&seq0))
	t.Log("addr: ", addr)
}
