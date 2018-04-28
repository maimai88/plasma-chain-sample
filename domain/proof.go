package domain

import "github.com/ethereum/go-ethereum/common"

type MerkleProof [MerkleProofSize]byte

func newMerkleProofFromBytes(b []byte) MerkleProof {
	proof := MerkleProof{}
	copy(proof[:], b[:])

	return proof
}

func (proof MerkleProof) Bytes() []byte {
	return proof[:]
}

func (proof MerkleProof) Hex() string {
	return common.ToHex(proof[:])
}
