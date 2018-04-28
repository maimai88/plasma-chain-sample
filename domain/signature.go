package domain

type Signature [SignatureLength]byte

func NewSignatureFromBytes(b []byte) Signature {
	sig := Signature{}
	copy(sig[:], b[:])

	if sig[SignatureLength-1] < SignatureRIRangeBase {
		sig[SignatureLength-1] += SignatureRIRangeBase
	}

	return sig
}

func (sig Signature) Bytes() []byte {
	return sig[:]
}
