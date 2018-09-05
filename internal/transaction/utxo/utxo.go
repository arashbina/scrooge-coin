package utxo

import (
	"encoding/binary"
	"hash/fnv"
)

// IsEqual returns true if two UTXOs are equal
// It infers equality by comparing hash and index values
func (u UTXO) IsEqual(u2 UTXO) bool {

	if (string(u.TxHash) == string(u2.TxHash)) && (u.Index == u2.Index) {
		return true
	}

	return false
}

// HashCode returns a unique hash value for a UTXO
// We can use this as key to a map in the UTXO pool
func (u UTXO) HashCode() uint64 {

	hasher := fnv.New64()
	b := make([]byte, 8)

	binary.LittleEndian.PutUint64(b, uint64(u.Index))
	hasher.Write(b)
	hasher.Write(u.TxHash)

	return hasher.Sum64()
}
