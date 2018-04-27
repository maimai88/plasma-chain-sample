package domain

type TxIn struct {
	BlockNum    uint
	TxIndex     uint
	OutputIndex uint
	Signature   Signature
}
