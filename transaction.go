//update v0.9
package dogecoin

import "bytes"

const (
	version  = "01000000"
	locktime = "00000000"
	sighash  = "01000000"
)

func CreateRawTransaction(unspent DogechainUnspent, balance DogechainBalance, dest *[]Destination) (string, uint64) {
	var rawtx bytes.Buffer
	var sendvalue uint64
	for i := 0; i < len(*dest); i++ {
		sendvalue += (*dest)[i].Value
	}
	inputstr, change := InputTemplate(sendvalue, unspent, balance, "")
	if inputstr == "saldo tidak mencukupi" {
		return inputstr, StrToInt(balance.Balance)
	} else {
		if change > 0 {
			(*dest) = append((*dest), Destination{unspent.UnspentOutputs[0].Address, change})
		}
		outputstr := OutputTemplate(*dest)
		rawtx.WriteString(version)
		rawtx.WriteString(inputstr)
		rawtx.WriteString(outputstr)
		rawtx.WriteString(locktime)
		rawtx.WriteString(sighash)
		return rawtx.String(), change
	}
}

func CreateSignedTransaction(unspent DogechainUnspent, balance DogechainBalance, dest []Destination, scriptsig string) (string, uint64) {
	var signedtx bytes.Buffer
	var sendvalue uint64
	for i := 0; i < (len(dest) - 1); i++ {
		sendvalue += dest[i].Value
	}
	inputstr, change := InputTemplate(sendvalue, unspent, balance, scriptsig)
	if inputstr == "saldo tidak mencukupi" {
		return inputstr, StrToInt(balance.Balance)
	} else {
		outputstr := OutputTemplate(dest)
		signedtx.WriteString(version)
		signedtx.WriteString(inputstr)
		signedtx.WriteString(outputstr)
		signedtx.WriteString(locktime)
		return signedtx.String(), change
	}
}
