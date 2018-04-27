package domain

import "github.com/ethereum/go-ethereum/common"

type Hash struct {
	common.Hash
}

func NewHashFromBytes(b []byte) *Hash {
	return &Hash{common.BytesToHash(b)}
}
