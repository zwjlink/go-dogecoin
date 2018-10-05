package dogecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/alivanz/go-crypto/bitcoin/base58"
)

const fee uint64 = 1000000
const sighashcode = "01"

type UnspentOutput struct {
	TxHash        string `json:"tx_hash"`
	TxOutputN     int    `json:"tx_output_n"`
	Script        string `json:"script"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
	Address       string `json:"address"`
}

type Destination struct {
	Address string
	Value   uint64
}

func OrderUnspent(unspent *DogechainUnspent) {
	var temp UnspentOutput
	for i := 1; i < len((*unspent).UnspentOutputs); i++ {
		for j := i; (j > 0) && (StrToInt((*unspent).UnspentOutputs[j].Value) > StrToInt((*unspent).UnspentOutputs[j-1].Value)); j-- {
			temp = (*unspent).UnspentOutputs[j]
			(*unspent).UnspentOutputs[j] = (*unspent).UnspentOutputs[j-1]
			(*unspent).UnspentOutputs[j-1] = temp
		}
	}
}

func InputTemplate(sendvalue uint64, unspent DogechainUnspent, scriptsig string) (string, uint64) {
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
		if scriptsig != "" {
			input.WriteString(VarInt(len(scriptsig) / 2))
			input.WriteString(scriptsig)
		} else {
			input.WriteString(VarInt(len(unspent.UnspentOutputs[i].Script) / 2))
			input.WriteString(unspent.UnspentOutputs[i].Script)
		}
		input.WriteString("ffffffff")
	}
	change := StrToInt(unspent.UnspentOutputs[i].Value) - sending
	input.WriteString(ReverseHex(unspent.UnspentOutputs[i].TxHash))
	index := fmt.Sprintf("%x", unspent.UnspentOutputs[i].TxOutputN)
	for len(index) < 8 {
		index = "0" + index
	}
	input.WriteString(ReverseHex(index))
	if scriptsig != "" {
		input.WriteString(VarInt(len(scriptsig) / 2))
		input.WriteString(scriptsig)
	} else {
		input.WriteString(VarInt(len(unspent.UnspentOutputs[i].Script) / 2))
		input.WriteString(unspent.UnspentOutputs[i].Script)
	}
	input.WriteString("ffffffff")
	i++
	inputfinal.WriteString(VarInt(i))
	inputfinal.WriteString(input.String())
	return inputfinal.String(), change
}

func OutputTemplate(dest []Destination) string {
	var output, outputfinal bytes.Buffer
	var i int
	for i = 0; i < len(dest); i++ {
		value := fmt.Sprintf("%x", dest[i].Value)
		for len(value) < 16 {
			value = "0" + value
		}
		output.WriteString(ReverseHex(value))
		binaddress, err := base58.Decode(dest[i].Address)
		ErrorCheck(err)
		pubkeyhash := BinAddressPubKeyHash(hex.EncodeToString(binaddress))
		scriptpubkey := P2PKH(pubkeyhash)
		output.WriteString(VarInt(len(scriptpubkey) / 2))
		output.WriteString(scriptpubkey)
	}
	outputfinal.WriteString(VarInt(i))
	outputfinal.WriteString(output.String())
	return outputfinal.String()
}

func ScriptSig(signature string, pubkey string) string {
	var scriptsig bytes.Buffer
	scriptsig.WriteString(VarInt((len(signature) / 2) + 1))
	scriptsig.WriteString(signature)
	scriptsig.WriteString(sighashcode)
	scriptsig.WriteString(VarInt(len(pubkey) / 2))
	return scriptsig.String()
}
