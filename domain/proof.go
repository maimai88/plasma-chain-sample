package domain

import "github.com/ethereum/go-ethereum/common"

type MerkleProof []byte

func (proof MerkleProof) Bytes() []byte {
	return []byte(proof)
}

func (proof MerkleProof) Hex() string {
	return common.ToHex(proof)
}
