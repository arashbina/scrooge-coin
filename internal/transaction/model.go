package transaction

type Transaction struct {
	Hash    []byte
	Inputs  []TXInput
	Outputs []TXOutput
}

type TXInput struct {
	PrevTxHash  []byte
	OutputIndex int
	Signature   []byte
}

type TXOutput struct {
	Value   float64
	Address []byte
}
