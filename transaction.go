package dogecoin

import "bytes"

const (
	version  = "01000000"
	locktime = "00000000"
	sighash  = "01000000"
)

func CreateRawTransaction(unspent DogechainUnspent, dest *[]Destination) string {
	var rawtx bytes.Buffer
	var sendvalue uint64
	for i := 0; i < len(*dest); i++ {
		sendvalue += (*dest)[i].Value
	}
	inputstr, change := InputTemplate(sendvalue, unspent)
	if change > 0 {
		(*dest) = append((*dest), Destination{unspent.UnspentOutputs[0].Address, change})
	}
	outputstr := OutputTemplate(*dest)
	rawtx.WriteString(version)
	rawtx.WriteString(inputstr)
	rawtx.WriteString(outputstr)
	rawtx.WriteString(locktime)
	rawtx.WriteString(sighash)
	return rawtx.String()
}
