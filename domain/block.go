package domain

import (
	"io"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	merkle "github.com/m0t0k1ch1/fixed-merkle"
)

type Block struct {
	Txes       []*Tx
	Signature  Signature
	merkleTree *merkle.Tree
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

func (block *Block) BuildMerkleTree() error {
	builder, err := merkle.NewTreeBuilder(sha3.NewKeccak256(), TxMerkleTreeDepth, TxMerkleLeafSize)
	if err != nil {
		return err
	}

	leaves := make([][]byte, len(block.Txes))
	for i, tx := range block.Txes {
		merkleHash, err := tx.MerkleHash()
		if err != nil {
			return err
		}
		leaves[i] = merkleHash.Bytes()
	}

	tree, err := builder.Build(leaves, true)
	if err != nil {
		return err
	}

	block.merkleTree = tree

	return nil
}

func (block *Block) MerkleRootHash() Hash {
	if block.merkleTree == nil {
		return Hash{}
	}

	return NewHashFromBytes(block.merkleTree.Root().Bytes())
}

func (block *Block) Sign(key *PrivateKey) error {
	blockHash, err := block.Hash()
	if err != nil {
		return err
	}

	sigBytes, err := key.Sign(blockHash.Bytes())
	if err != nil {
		return err
	}

	block.Signature = NewSignatureFromBytes(sigBytes)

	return nil
}
