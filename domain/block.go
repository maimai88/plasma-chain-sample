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

	return newHashFromBytes(crypto.Keccak256(b)), nil
}

func (blk *Block) BuildMerkleTree() error {
	conf, err := merkle.NewConfig(
		sha3.NewKeccak256(),
		TxMerkleTreeDepth,
		TxMerkleLeafSize,
	)
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

	tree, err := merkle.NewTree(conf, leaves, true)
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

	return newHashFromBytes(blk.merkleTree.Root().Bytes())
}

func (blk *Block) CreateMerkleProof(txIndex int) (MerkleProof, error) {
	if err := blk.validateTxIndex(txIndex); err != nil {
		return MerkleProof{}, err
	}

	b, err := blk.merkleTree.CreateMembershipProof(txIndex)
	if err != nil {
		return MerkleProof{}, err
	}

	return newMerkleProofFromBytes(b), err
}

func (blk *Block) VerifyMerkleProof(txIndex int, proof MerkleProof) (bool, error) {
	if err := blk.validateTxIndex(txIndex); err != nil {
		return false, err
	}

	return blk.merkleTree.VerifyMembershipProof(txIndex, proof.Bytes())
}

func (blk *Block) Sign(privkey *PrivateKey) error {
	blkHash, err := blk.Hash()
	if err != nil {
		return err
	}

	sig, err := privkey.Sign(blkHash.Bytes())
	if err != nil {
		return err
	}

	blk.Signature = sig

	return nil
}

func (blk *Block) validateTxIndex(txIndex int) error {
	if txIndex < 0 || len(blk.Txes) <= txIndex {
		return ErrTxIndexOutOfRange
	}

	return nil
}
