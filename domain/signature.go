package domain

const (
	SignatureLength = 65
)

type Signature [SignatureLength]byte

func NewSignatureFromBytes(b []byte) Signature {
	sig := Signature{}
	copy(sig[:], b[:])

	return sig
}

func (sig Signature) Bytes() []byte {
	return sig[:]
}
