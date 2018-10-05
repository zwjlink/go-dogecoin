package dogecoin

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestGetUnspentData(t *testing.T) {
	var signature, pubkey, scriptsig string
	unspent := GetUnspent("DMr3fEiVrPWFpoCWS958zNtqgnFb7QWn9D")
	OrderUnspent(&unspent)
	for _, row := range unspent.UnspentOutputs {
		log.Printf("%v %v %v", row.TxHash, row.TxOutputN, row.Value)
	}
	dest := make([]Destination, 0)
	dest = append(dest, Destination{"DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi", 2570000000000})
	dest = append(dest, Destination{"DRsLzWNubViGtLoxXnaNaMbvTiadMMsUYL", 1250000000000})
	rawtransaction := CreateRawTransaction(unspent, &dest)
	log.Println(rawtransaction)
	rawtxbyte, err := hex.DecodeString(rawtransaction)
	ErrorCheck(err)
	rawtxhash := Hash(rawtxbyte)
	log.Println(hex.EncodeToString(rawtxhash))
	/*signing process*/
	//add here
	/*create signed transaction*/
	ScriptSig(signature, pubkey)
	signtx, change := CreateSignedTransaction(unspent, dest, scriptsig)
	log.Println(signtx)
	log.Println(change)
}
