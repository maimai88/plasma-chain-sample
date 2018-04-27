package domain

import "bytes"

type Signature [SignatureLength]byte

func NewSignatureFromBytes(b []byte) Signature {
	sig := Signature{}
	copy(sig[:], b[:])

	return sig
}

func (sig Signature) Bytes() []byte {
	return sig[:]
}

func (sig *Signature) SwitchToHigherRIRange() {
	if bytes.Equal(sig.Bytes(), Signature{}.Bytes()) {
		return
	}

	if sig[SignatureLength-1] < SignatureRIRangeBase {
		sig[SignatureLength-1] += SignatureRIRangeBase
	}
}
