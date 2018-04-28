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
func (blk *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{blk.Txes, blk.Signature})
}

func (blk *Block) Hash() (Hash, error) {
	b, err := rlp.EncodeToBytes([]interface{}{blk.Txes})
	if err != nil {
		return Hash{}, nil
	}

	return NewHashFromBytes(crypto.Keccak256(b)), nil
}

func (blk *Block) BuildMerkleTree() error {
	builder, err := merkle.NewTreeBuilder(sha3.NewKeccak256(), TxMerkleTreeDepth, TxMerkleLeafSize)
	if err != nil {
		return err
	}

	leaves := make([][]byte, len(blk.Txes))
	for i, tx := range blk.Txes {
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

	blk.merkleTree = tree

	return nil
}

func (blk *Block) MerkleRootHash() Hash {
	if blk.merkleTree == nil {
		return Hash{}
	}

	return NewHashFromBytes(blk.merkleTree.Root().Bytes())
}

func (blk *Block) CreateMerkleProof(txIndex int) (MerkleProof, error) {
	if txIndex < 0 || len(blk.Txes) <= txIndex {
		return nil, ErrTxIndexOutOfRange
	}

	b, err := blk.merkleTree.CreateMembershipProof(txIndex)
	if err != nil {
		return nil, err
	}

	return MerkleProof(b), err
}

func (blk *Block) Sign(key *PrivateKey) error {
	blkHash, err := blk.Hash()
	if err != nil {
		return err
	}

	sig, err := key.Sign(blkHash.Bytes())
	if err != nil {
		return err
	}

	blk.Signature = sig

	return nil
}
