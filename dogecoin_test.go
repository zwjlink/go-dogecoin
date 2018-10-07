//update v0.9
package dogecoin

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestGetUnspentData(t *testing.T) {
	// var pubkey, scriptsig string
	Address := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	unspent := GetUnspent(Address)
	balance := GetBalance(Address)
	OrderUnspent(&unspent)
	// for _, row := range unspent.UnspentOutputs {
	// 	log.Printf("%v %v %v", row.TxHash, row.TxOutputN, row.Value)
	// }
	dest := make([]Destination, 0)
	sendvalue := uint64(6300000000)
	log.Printf("sendvalue : %v\n", sendvalue)
	dest = append(dest, Destination{"DBHdUWoMQAAt4CgCfWztwRovxhQQ9qo5Um", sendvalue})
	rawtransaction, change := CreateRawTransaction(unspent, balance, &dest)
	log.Printf("rawtxhex  : %v\n", rawtransaction)
	if rawtransaction == "saldo tidak mencukupi" {
		log.Printf("balance   : %v\n", change)
	} else {
		log.Printf("change    : %v\n", change)
		rawtxbyte, err := hex.DecodeString(rawtransaction)
		ErrorCheck(err)
		rawtxhash := Hash(rawtxbyte)
		log.Printf("rawtxhash : %v", hex.EncodeToString(rawtxhash))
	}
	// r, s, err := ecdsa.Sign(rand.Reader, /*privatekey*/, rawtxhash)
	// ErrorCheck(err)
	// ScriptSig(EvenCorrect(r), EvenCorrect(s), /*pubkey*/)
	// signtx, change := CreateSignedTransaction(unspent, dest, scriptsig)
	// log.Println(signtx)
}
