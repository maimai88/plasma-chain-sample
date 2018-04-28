package domain

import "github.com/ethereum/go-ethereum/common"

type Address common.Address

func NewAddressFromHex(addressHex string) Address {
	return Address(common.HexToAddress(addressHex))
}

func (address Address) Bytes() []byte {
	return common.Address(address).Bytes()
}

func (address Address) Hex() string {
	return common.Address(address).Hex()
}
