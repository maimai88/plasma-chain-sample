package domain

import "github.com/ethereum/go-ethereum/common"

type Hash common.Hash

func newHashFromBytes(hashBytes []byte) Hash {
	return Hash(common.BytesToHash(hashBytes))
}

func (h Hash) Bytes() []byte {
	return common.Hash(h).Bytes()
}

func (h Hash) Hex() string {
	return common.Hash(h).Hex()
}
