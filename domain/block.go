package domain

import (
	"io"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Block struct {
	Txes      []*Tx
	Signature Signature
}

func NewBlock(txes []*Tx) *Block {
	return &Block{
		Txes: txes,
	}
}

// implements RLP Encoder interface.
//
// ref. https://godoc.org/github.com/ethereum/go-ethereum/rlp#Encoder
func (block *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{block.Txes, block.Signature})
}

func (block *Block) Hash() (Hash, error) {
	b, err := rlp.EncodeToBytes([]interface{}{block.Txes})
	if err != nil {
		return Hash{}, nil
	}

	return NewHashFromBytes(crypto.Keccak256(b)), nil
}
