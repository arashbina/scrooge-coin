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

	input := t.Inputs[index]

	err := input.addBytesNoSig(&b)
	if err != nil {
		return nil, err
	}

	// convert all outputs to bytes and put into the buffer
	for _, output := range t.Outputs {

		err := output.addBytes(&b)
		if err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

// GetRawTransaction returns raw transaction to be used when signing a complete block
func (t Transaction) GetRawTransaction() ([]byte, error) {

	// create a buffer to put bytes in as we convert various components of the transaction
	var b bytes.Buffer

	for _, input := range t.Inputs {

		err := input.addBytesNoSig(&b)
		if err != nil {
			return nil, err
		}

		_, err = b.Write(input.Signature)
		if err != nil {
			return nil, err
		}
	}

	for _, output := range t.Outputs {

		err := output.addBytes(&b)
		if err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (input TXInput) addBytesNoSig(buf *bytes.Buffer) error {

	// add previous hash
	_, err := buf.Write(input.PrevTxHash)
	if err != nil {
		return err
	}

	// add corresponding output index
	bi := make([]byte, 8)
	binary.LittleEndian.PutUint64(bi, uint64(input.OutputIndex))
	_, err = buf.Write(bi)
	if err != nil {
		return err
	}

	return nil
}

func (output TXOutput) addBytes(buf *bytes.Buffer) error {

	// add output value
	bv := make([]byte, 8)
	binary.LittleEndian.PutUint64(bv, math.Float64bits(output.Value))
	_, err := buf.Write(bv)
	if err != nil {
		return err
	}

	// add address/public key
	_, err = buf.Write(output.Address)
	if err != nil {
		return err
	}

	return nil
}
