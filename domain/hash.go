package domain

import "github.com/ethereum/go-ethereum/common"

type Hash common.Hash

func NewHashFromBytes(b []byte) Hash {
	return Hash(common.BytesToHash(b))
}

func (h Hash) Bytes() []byte {
	return common.Hash(h).Bytes()
}

func (h Hash) Hex() string {
	return common.Hash(h).Hex()
}
