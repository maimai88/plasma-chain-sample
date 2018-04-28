package domain

import "github.com/ethereum/go-ethereum/common"

type MerkleProof [MerkleProofSize]byte

func newMerkleProofFromBytes(proofBytes []byte) MerkleProof {
	proof := MerkleProof{}
	copy(proof[:], proofBytes[:])

	return proof
}

func (proof MerkleProof) Bytes() []byte {
	return proof[:]
}

func (proof MerkleProof) Hex() string {
	return common.ToHex(proof[:])
}
