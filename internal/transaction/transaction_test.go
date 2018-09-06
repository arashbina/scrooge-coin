package transaction

import (
	"bytes"
	"testing"
)

func TestGetRawDataToSign(t *testing.T) {

	input := TXInput{
		PrevTxHash:  []byte("hash1"),
		OutputIndex: 0,
		Signature:   []byte("hash1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
	}

	output := TXOutput{
		Value:   1.0,
		Address: []byte("publicKey1"),
	}

	tx := Transaction{
		Inputs:  []TXInput{input},
		Outputs: []TXOutput{output},
	}

	data, err := tx.GetRawDataToSign(0)
	if data == nil || err != nil {
		t.Errorf("error getting data to sign: %s", err)
	}

	if cmp := bytes.Compare(input.Signature, data); cmp != 0 {
		t.Errorf("signature mismatch. Expected: %q, Got: %q", input.Signature, data)
	}
}

func TestGetRawTransaction(t *testing.T) {

	input := TXInput{
		PrevTxHash:  []byte("hash1"),
		OutputIndex: 0,
		Signature:   []byte("hash1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?publicKey1"),
	}

	output := TXOutput{
		Value:   1.0,
		Address: []byte("publicKey1"),
	}

	tx := Transaction{
		Inputs:  []TXInput{input},
		Outputs: []TXOutput{output},
	}

	data, err := tx.GetRawTransaction()
	if err != nil || data == nil {
		t.Errorf("could not get transaction: %s", err)
	}
}
