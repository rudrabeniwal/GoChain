package main

import "fmt"

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

const subsidy = 10

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

type TXInput struct {
	Txid      []byte
	Vout      int    //stores an index of an output in the transaction
	ScriptSig string //ScriptSig is a script which provides data to be used in an output’s ScriptPubKey
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to} //subsidy is the amount of reward
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetId()

	return &tx
}
