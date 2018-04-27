package domain

import (
	"io"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Tx struct {
	Inputs  [2]*TxIn
	Outputs [2]*TxOut
}

type TxIn struct {
	BlockNum    uint
	TxIndex     uint
	OutputIndex uint
	Signature   Signature
}

type TxOut struct {
	Address Address
	Amount  uint
}

func NewTx() *Tx {
	return &Tx{}
}

func (tx *Tx) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		tx.Inputs[0].BlockNum, tx.Inputs[0].TxIndex, tx.Inputs[0].OutputIndex,
		tx.Inputs[1].BlockNum, tx.Inputs[1].TxIndex, tx.Inputs[1].OutputIndex,
		tx.Outputs[0].Address, tx.Outputs[0].Amount,
		tx.Outputs[1].Address, tx.Outputs[1].Amount,
	})
}

func (tx *Tx) SetTxIn(index uint, blockNum uint, txIndex uint, outputIndex uint) {
	tx.Inputs[index] = &TxIn{
		BlockNum:    blockNum,
		TxIndex:     txIndex,
		OutputIndex: outputIndex,
		Signature:   Signature{},
	}
}

func (tx *Tx) SetTxOut(index uint, address Address, amount uint) {
	tx.Outputs[index] = &TxOut{
		Address: address,
		Amount:  amount,
	}
}

func (tx *Tx) Hash() (Hash, error) {
	b, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return Hash{}, err
	}

	return NewHashFromBytes(crypto.Keccak256(b)), nil
}

func (tx *Tx) Sign(inputIndex uint, key *PrivateKey) error {
	txHash, err := tx.Hash()
	if err != nil {
		return err
	}

	sigBytes, err := key.Sign(txHash.Bytes())
	if err != nil {
		return err
	}

	tx.Inputs[inputIndex].Signature = NewSignatureFromBytes(sigBytes)

	return nil
}