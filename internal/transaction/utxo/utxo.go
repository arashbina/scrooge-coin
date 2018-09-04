package utxo

func (u UTXO) IsEqual(u2 UTXO) bool {

	if (string(u.TxHash) == string(u2.TxHash)) && (u.Index == u2.Index) {
		return true
	}

	return false
}
