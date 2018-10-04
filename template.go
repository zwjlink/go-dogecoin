package dogecoin

import (
	"bytes"
	"fmt"
)

const fee uint64 = 1000000

//type InputByte []byte
type UnspentOutput struct {
	TxHash        string `json:"tx_hash"`
	TxOutputN     int    `json:"tx_output_n"`
	Script        string `json:"script"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
	Address       string `json:"address"`
}

func OrderUnspent(unspent *DogechainUnspent) {
	var temp UnspentOutput
	for i := 1; i < len((*unspent).UnspentOutputs); i++ {
		for j := i; (j > 0) && (StrToInt((*unspent).UnspentOutputs[j].Value) > StrToInt((*unspent).UnspentOutputs[j-1].Value)); j-- {
			//swap
			temp = (*unspent).UnspentOutputs[j]
			(*unspent).UnspentOutputs[j] = (*unspent).UnspentOutputs[j-1]
			(*unspent).UnspentOutputs[j-1] = temp
		}
	}
}

func InputTemplate(sendvalue uint64, unspent DogechainUnspent) (string, uint64) {
	var input, inputfinal bytes.Buffer
	var i int
	sending := sendvalue + fee
	for i = 0; sending > StrToInt(unspent.UnspentOutputs[i].Value); i++ {
		sending = sending - StrToInt(unspent.UnspentOutputs[i].Value)
		input.WriteString(ReverseHex(unspent.UnspentOutputs[i].TxHash))
		index := fmt.Sprintf("%x", unspent.UnspentOutputs[i].TxOutputN)
		for len(index) < 8 {
			index = "0" + index
		}
		input.WriteString(ReverseHex(index))
		input.WriteString(fmt.Sprintf("%x", (len(unspent.UnspentOutputs[i].Script) / 2)))
		input.WriteString(unspent.UnspentOutputs[i].Script)
		input.WriteString("ffffffff")
	}
	change := StrToInt(unspent.UnspentOutputs[i].Value) - sending
	input.WriteString(ReverseHex(unspent.UnspentOutputs[i].TxHash))
	index := fmt.Sprintf("%x", unspent.UnspentOutputs[i].TxOutputN)
	for len(index) < 8 {
		index = "0" + index
	}
	input.WriteString(ReverseHex(index))
	input.WriteString(fmt.Sprintf("%x", (len(unspent.UnspentOutputs[i].Script) / 2)))
	input.WriteString(unspent.UnspentOutputs[i].Script)
	input.WriteString("ffffffff")
	if (len(fmt.Sprint(i)) % 2) > 0 {
		inputfinal.WriteString("0" + fmt.Sprintf("%x", i))
	} else {
		inputfinal.WriteString(fmt.Sprintf("%x", i))
	}
	inputfinal.WriteString(input.String())
	return inputfinal.String(), change
}
