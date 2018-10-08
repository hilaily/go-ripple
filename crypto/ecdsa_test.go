package crypto

import (
	"fmt"
	"testing"
)

func TestGenKey(t *testing.T) {
	seed := []byte("snh1zUj8AKjdLPDRapFGpJeaBRDHm")
	key := NewECDSAKey(seed)
	fmt.Printf("pub: ", key.PubKey().SerializeCompressed())
}
