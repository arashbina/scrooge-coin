package utxo

import (
	"testing"
)

func TestCreateDeleteUTXO(t *testing.T) {

	// reset pool since we are not sure in what order tests run
	pool = make(UTXOPool)

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

func TestNewPoolFromPool(t *testing.T) {

	// reset pool since we are not sure in what order tests run
	pool = make(UTXOPool)

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

func TestGetAllUTXOs(t *testing.T) {

	// reset pool since we are not sure in what order tests run
	pool = make(UTXOPool)

	u := UTXO{
		TxHash: []byte("hash1"),
		Index:  1,
	}

	AddUTXO(u)

	u.TxHash = []byte("hash2")
	u.Index = 2
	AddUTXO(u)

	utxos := GetAllUTXOs()
	if len(utxos) != 2 {
		t.Errorf("expected 2 UTXOs but got: %d", len(utxos))
	}
}

func TestPoolContainsUTXO(t *testing.T) {

	// reset pool since we are not sure in what order tests run
	pool = make(UTXOPool)

	u := UTXO{
		TxHash: []byte("hash1"),
		Index:  1,
	}

	if PoolContainsUTXO(u) {
		t.Errorf("did not expect the main pool to contain the utxo but it did")
	}

	AddUTXO(u)

	if !PoolContainsUTXO(u) {
		t.Errorf("expected the main pool to contain the utxo but it didn't")
	}
}
