package transaction

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func (t Transaction) GetRawDataToSign(index int) ([]byte, error) {

	var b bytes.Buffer

	if index >= len(t.Inputs) {
		return nil, fmt.Errorf("requested input index does not exist in the transation")
	}

	input := t.Inputs[index]
	_, err := b.Write(input.PrevTxHash)
	if err != nil {
		return nil, err
	}

	bi := make([]byte, 8)
	binary.LittleEndian.PutUint64(bi, uint64(index))
	_, err = b.Write(bi)
	if err != nil {
		return nil, err
	}

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

	rawBytes := make([]byte, b.Len())
	_, err = b.Read(rawBytes)
	if err != nil {
		return nil, err
	}

	return rawBytes, nil
}
