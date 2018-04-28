package domain

import "github.com/ethereum/go-ethereum/common"

type Signature [SignatureSize]byte

func newSignatureFromBytes(b []byte) Signature {
	sig := Signature{}
	copy(sig[:], b[:])

	if sig[SignatureSize-1] < SignatureRIRangeBase {
		sig[SignatureSize-1] += SignatureRIRangeBase
	}

	return sig
}

func (sig Signature) Bytes() []byte {
	return sig[:]
}

func (sig Signature) Hex() string {
	return common.ToHex(sig[:])
}
