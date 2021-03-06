package utxo

type UTXOPool map[uint64]UTXO

// pool keeps a map of all the UTXOs.
// Since we only need one pool the pool is private
var pool = make(UTXOPool)

// AddUTXO adds a utxo to the UTXOPool
func AddUTXO(u UTXO) {

	hash := u.HashCode()
	pool[hash] = u
}

// RemoveUTXO removes a utxo from the UTXOPool
func RemoveUTXO(u UTXO) {

	hash := u.HashCode()
	delete(pool, hash)
}

// NewPoolFromPool reinitializes the main pool and copies the elements of the
// supplied pool to the main pool
func NewPoolFromPool(p UTXOPool) {

	pool = make(UTXOPool)
	for k, v := range p {
		pool[k] = v
	}
}

// GetAllUTXOs returns all UTXOs in the main pool
func GetAllUTXOs() []UTXO {

	utxos := make([]UTXO, len(pool))

	i := 0
	for _, u := range pool {

		utxos[i] = u
		i++
	}

	return utxos
}

func PoolContainsUTXO(u UTXO) bool {

	h := u.HashCode()
	_, ok := pool[h]

	return ok
}
