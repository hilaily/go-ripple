package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenKey(t *testing.T) {
	secret := "snh1zUj8AKjdLPDRapFGpJeaBRDHm"
	seed, err := NewRippleHash(secret)
	fmt.Printf("%x\n", seed.Payload())
	key, err := NewECDSAKey(seed.Payload())
	var sequenceZero uint32
	h, err := AccountPrivateKey(key, &sequenceZero)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("hash: ", h.String())
	h, _ = AccountPublicKey(key, &sequenceZero)
	fmt.Println("hash: ", h.String())
	fmt.Printf("pri: %x\n", key.Private(&sequenceZero))
	fmt.Printf("pub: %x\n", key.Public(&sequenceZero))

	priKeyStr := "86029426A6D950A14CEDD1AE33F0EB8C7CE1C0E8190D41D82C52EA160084B9E8"
	bt, _ := hex.DecodeString(priKeyStr)
	key1 := LoadECDSKey(bt)
	fmt.Printf("pub: %x\n", key1.Private(nil))
}
