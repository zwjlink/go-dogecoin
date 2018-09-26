package dogecoin

import (
	"log"
	"testing"
)

func TestGetAddressData(t *testing.T) {
	//proxy added later
	balance := GetBalance("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(balance.Balance)
	log.Println(DogeValueStrToInt(balance.Balance))
	received := GetReceived("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(received.Received)
	log.Println(DogeValueStrToInt(received.Received))
	sent := GetSent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(sent.Sent)
	log.Println(DogeValueStrToInt(sent.Sent))
	unspent := GetUnspent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(unspent.UnspentOutputs)
	for _, row := range unspent.UnspentOutputs {
		log.Printf("%v %v", row.TxHash, row.TxOutputN)
	}
	log.Println("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
}
