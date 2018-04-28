package domain

import (
	"bytes"
	"io"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Tx struct {
	Inputs  [2]*TxIn
	Outputs [2]*TxOut
}

func NewTx() *Tx {
	return &Tx{}
}

// implements RLP Encoder interface.
//
// ref. https://godoc.org/github.com/ethereum/go-ethereum/rlp#Encoder
func (tx *Tx) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		tx.Inputs[0].BlockNum, tx.Inputs[0].TxIndex, tx.Inputs[0].OutputIndex,
		tx.Inputs[1].BlockNum, tx.Inputs[1].TxIndex, tx.Inputs[1].OutputIndex,
		tx.Outputs[0].Address.Bytes(), tx.Outputs[0].Amount,
		tx.Outputs[1].Address.Bytes(), tx.Outputs[1].Amount,
		tx.Inputs[0].Signature.Bytes(),
		tx.Inputs[1].Signature.Bytes(),
	})
}

func (tx *Tx) SetTxIn(index uint, blkNum uint, txIndex uint, outputIndex uint) {
	tx.Inputs[index] = &TxIn{
		BlockNum:    blkNum,
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
	b, err := rlp.EncodeToBytes([]interface{}{
		tx.Inputs[0].BlockNum, tx.Inputs[0].TxIndex, tx.Inputs[0].OutputIndex,
		tx.Inputs[1].BlockNum, tx.Inputs[1].TxIndex, tx.Inputs[1].OutputIndex,
		tx.Outputs[0].Address.Bytes(), tx.Outputs[0].Amount,
		tx.Outputs[1].Address.Bytes(), tx.Outputs[1].Amount,
	})
	if err != nil {
		return Hash{}, err
	}

	return newHashFromBytes(crypto.Keccak256(b)), nil
}

func (tx *Tx) MerkleHash() (Hash, error) {
	txHash, err := tx.Hash()
	if err != nil {
		return Hash{}, err
	}

	buf := bytes.NewBuffer(txHash.Bytes())
	if _, err := buf.Write(tx.Inputs[0].Signature.Bytes()); err != nil {
		return Hash{}, err
	}
	if _, err := buf.Write(tx.Inputs[1].Signature.Bytes()); err != nil {
		return Hash{}, err
	}

	return newHashFromBytes(crypto.Keccak256(buf.Bytes())), nil
}

func (tx *Tx) Sign(inputIndex uint, key *PrivateKey) error {
	txHash, err := tx.Hash()
	if err != nil {
		return err
	}

	sig, err := key.Sign(txHash.Bytes())
	if err != nil {
		return err
	}

	tx.Inputs[inputIndex].Signature = sig

	return nil
}
