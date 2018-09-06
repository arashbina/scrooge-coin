package transaction

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// GetRawDataToSign returns the data needed to be signed for an input of a transaction
func (t Transaction) GetRawDataToSign(index int) ([]byte, error) {

	// create a buffer to put bytes in as we convert various components of the transaction
	var b bytes.Buffer

	// check to make sure the requested index in within the []TXInput range
	if index >= len(t.Inputs) {
		return nil, fmt.Errorf("requested input index does not exist in the transation")
	}

	// put in the previous hash of the input into the buffer
	input := t.Inputs[index]
	_, err := b.Write(input.PrevTxHash)
	if err != nil {
		return nil, err
	}

	// convert the index into bytes and put into the buffer
	bi := make([]byte, 8)
	binary.LittleEndian.PutUint64(bi, uint64(index))
	_, err = b.Write(bi)
	if err != nil {
		return nil, err
	}

	// convert all outputs to bytes and put into the buffer
	for _, output := range t.Outputs {
		bv := make([]byte, 8)
		binary.LittleEndian.PutUint64(bv, math.Float64bits(output.Value))
		_, err = b.Write(bv)
		if err != nil {
			return nil, err
		}
		_, err = b.Write(output.Address)
		if err != nil {
			return nil, err
		}
	}

	// read the bytes off the buffer and send them back
	rawBytes := make([]byte, b.Len())
	_, err = b.Read(rawBytes)
	if err != nil {
		return nil, err
	}

	return rawBytes, nil
}

// GetRawTransaction returns raw transaction to be used when signing a complete block
func (t Transaction) GetRawTransaction() ([]byte, error) {

	return nil, nil
}
