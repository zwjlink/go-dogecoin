//update v0.9
package dogecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	crypto "github.com/alivanz/go-crypto"
	"github.com/alivanz/go-crypto/bitcoin/base58"
)

const (
	fee         uint64 = 100000000
	sighashcode        = "01"
	header             = "30"
	integer            = "02"
)

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

func CanSpent(balance DogechainBalance, sendvalue uint64) string {
	if StrToInt(balance.Balance) < (sendvalue + fee) {
		return "saldo tidak mencukupi"
	} else {
		return ""
	}
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

func ChangeUnspent(sendvalue uint64, unspent *DogechainUnspent, dest *[]Destination) int {
	var i int
	sending := sendvalue + fee
	for i = 0; sending > 0; i++ {
		if sending >= StrToInt((*unspent).UnspentOutputs[i].Value) {
			sending = sending - StrToInt((*unspent).UnspentOutputs[i].Value)
		} else {
			change := StrToInt((*unspent).UnspentOutputs[i].Value) - sending
			(*dest) = append((*dest), Destination{(*unspent).UnspentOutputs[0].Address, change})
		}
	}
	return i
}

func InputTemplate(unspent DogechainUnspent, dest []Destination, wallet crypto.Wallet, numindex int, posindex int) string {
	var input, inputfinal bytes.Buffer
	var i int
	var scriptsig, index string
	for i = 0; i < numindex; i++ {
		input.WriteString(ReverseHex(unspent.UnspentOutputs[i].TxHash))
		index = fmt.Sprintf("%x", unspent.UnspentOutputs[i].TxOutputN)
		for len(index) < 8 {
			index = "0" + index
		}
		input.WriteString(ReverseHex(index))
		switch {
		case wallet != nil:
			scriptsig = CreateSignature(unspent, dest, wallet, numindex, i)
			input.WriteString(VarInt(len(scriptsig) / 2))
			input.WriteString(scriptsig)
		case (wallet == nil) && (i == posindex):
			input.WriteString(VarInt(len(unspent.UnspentOutputs[i].Script) / 2))
			input.WriteString(unspent.UnspentOutputs[i].Script)
		default:
			input.WriteString("00")
		}
		input.WriteString("ffffffff")
	}
	inputfinal.WriteString(VarInt(i))
	inputfinal.WriteString(input.String())
	return inputfinal.String()
}

func OutputTemplate(dest []Destination) string {
	var output, outputfinal bytes.Buffer
	var i int
	var value, pubkeyhash, scriptpubkey string
	for i = 0; i < len(dest); i++ {
		value = fmt.Sprintf("%x", dest[i].Value)
		for len(value) < 16 {
			value = "0" + value
		}
		output.WriteString(ReverseHex(value))
		binaddress, err := base58.Decode(dest[i].Address)
		ErrorCheck(err)
		pubkeyhash = BinAddressPubKeyHash(hex.EncodeToString(binaddress))
		scriptpubkey = P2PKH(pubkeyhash)
		output.WriteString(VarInt(len(scriptpubkey) / 2))
		output.WriteString(scriptpubkey)
	}
	outputfinal.WriteString(VarInt(i))
	outputfinal.WriteString(output.String())
	return outputfinal.String()
}

func ScriptSig(r, s, pubkey string) string {
	var sign, signfinal, scriptsig bytes.Buffer

	sign.WriteString(integer)
	sign.WriteString(VarInt(len(r) / 2))
	sign.WriteString(r)
	sign.WriteString(integer)
	sign.WriteString(VarInt(len(s) / 2))
	sign.WriteString(s)

	signfinal.WriteString(header)
	signfinal.WriteString(VarInt(len(sign.String()) / 2))
	signfinal.WriteString(sign.String())
	signfinal.WriteString(sighashcode)

	scriptsig.WriteString(VarInt(len(signfinal.String()) / 2))
	scriptsig.WriteString(signfinal.String())
	scriptsig.WriteString(VarInt(len(pubkey) / 2))
	scriptsig.WriteString(pubkey)

	return scriptsig.String()
}
