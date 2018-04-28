package domain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PrivateKey struct {
	*ecdsa.PrivateKey
}

func NewPrivateKeyFromHex(keyHex string) (*PrivateKey, error) {
	privkey, err := crypto.ToECDSA(common.FromHex(keyHex))
	if err != nil {
		return nil, err
	}

	return &PrivateKey{privkey}, nil
}

func (key *PrivateKey) Sign(b []byte) (Signature, error) {
	b, err := crypto.Sign(b, key.PrivateKey)
	if err != nil {
		return Signature{}, err
	}

	return NewSignatureFromBytes(b), nil
}
