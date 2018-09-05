package utxo

import "testing"

func TestEquality(t *testing.T) {

	tests := []struct {
		u1  UTXO
		u2  UTXO
		res bool
	}{
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			res: true,
		},
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u2"),
				Index:  2,
			},
			res: false,
		},
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u1"),
				Index:  2,
			},
			res: false,
		},
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u2"),
				Index:  1,
			},
			res: false,
		},
	}

	for i, test := range tests {
		if test.u1.IsEqual(test.u2) != test.res {
			t.Errorf("test: %d\n expected the comparison to be %t but wasn't", i, test.res)
		}
	}
}

func TestHashCode(t *testing.T) {

	tests := []struct {
		u1  UTXO
		u2  UTXO
		res bool
	}{
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			res: true,
		},
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u2"),
				Index:  2,
			},
			res: false,
		},
		{
			u1: UTXO{
				TxHash: []byte("u1"),
				Index:  1,
			},
			u2: UTXO{
				TxHash: []byte("u2"),
				Index:  1,
			},
			res: false,
		},
	}

	for i, test := range tests {
		if (test.u1.HashCode() == test.u2.HashCode()) != test.res {
			t.Errorf("test: %d\n expected hash codes to be %t but wasn't", i, test.res)
		}
	}
}
