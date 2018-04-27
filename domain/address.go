package domain

import "github.com/ethereum/go-ethereum/common"

type Address struct {
	common.Address
}

func NewAddressFromHex(addressHex string) *Address {
	return &Address{common.HexToAddress(addressHex)}
}
