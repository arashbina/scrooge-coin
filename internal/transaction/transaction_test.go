package transaction

import (
	"bytes"
	"testing"
)

func TestGetRawDataToSign(t *testing.T) {

	tests := []struct {
		tx       Transaction
		expected []byte
		res      bool
	}{
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      true,
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 2,
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //outputIndex mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash2"),
						OutputIndex: 1,
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //prevTxHash mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.2,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //Value mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey2"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //Address mismatch
		},
	}

	for i, test := range tests {
		data, err := test.tx.GetRawDataToSign(0)
		if data == nil || err != nil {
			t.Errorf("test: %d\n error getting data to sign: %s", i, err)
		}

		if cmp := bytes.Compare(test.expected, data); (cmp == 0) != test.res {
			t.Errorf("test: %d\n signature mismatch. Expected: %q, Got: %q", i, test.expected, data)
		}
	}

}

func TestGetRawTransaction(t *testing.T) {
	tests := []struct {
		tx       Transaction
		expected []byte
		res      bool
	}{
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
						Signature:   []byte("signature1"),
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00signature1\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      true,
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash2"),
						OutputIndex: 1,
						Signature:   []byte("signature1"),
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00signature1\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //hash mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 2,
						Signature:   []byte("signature1"),
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00signature1\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //outputIndex mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
						Signature:   []byte("signature2"),
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00signature1\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //signature mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
						Signature:   []byte("signature1"),
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.2,
						Address: []byte("publicKey1"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00signature1\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //value mismatch
		},
		{
			tx: Transaction{
				Inputs: []TXInput{
					{
						PrevTxHash:  []byte("hash1"),
						OutputIndex: 1,
						Signature:   []byte("signature1"),
					},
				},
				Outputs: []TXOutput{
					{
						Value:   1.0,
						Address: []byte("publicKey2"),
					},
				},
			},
			expected: []byte("hash1\x01\x00\x00\x00\x00\x00\x00\x00signature1\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
			res:      false, //address mismatch
		},
	}

	for i, test := range tests {
		data, err := test.tx.GetRawTransaction()
		if data == nil || err != nil {
			t.Errorf("test: %d\n could not get transaction: %s", i, err)
		}

		if cmp := bytes.Compare(test.expected, data); (cmp == 0) != test.res {
			t.Errorf("test: %d\n bytes mismatch. Expected: %q, Got: %q", i, test.expected, data)
		}
	}
}
