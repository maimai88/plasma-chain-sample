package domain

type TxOut struct {
	Address Address
	Amount  uint
	IsSpent bool
}
