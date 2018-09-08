package utxo

import (
	"testing"
)

func TestCreateDeleteUTXO(t *testing.T) {

	u := UTXO{
		TxHash: []byte("hash"),
		Index:  1,
	}

	AddUTXO(u)
	if _, ok := pool[u.HashCode()]; !ok {
		t.Errorf("expected utxo to exit in the pool but didn't")
	}

	RemoveUTXO(u)
	if _, ok := pool[u.HashCode()]; ok {
		t.Errorf("expected utxo to not exit but it did")
	}
}

func TestCreatePoolFromPool(t *testing.T) {

	u := UTXO{
		TxHash: []byte("hash"),
		Index:  1,
	}

	h := u.HashCode()
	p := make(UTXOPool)
	p[h] = u

	NewPoolFromPool(p)
	if _, ok := pool[u.HashCode()]; !ok {
		t.Errorf("expected utxo to exit in the pool but didn't")
	}
}
