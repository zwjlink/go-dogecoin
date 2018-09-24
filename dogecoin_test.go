package dogecoin

import (
	"log"
	"testing"
)

func TestGetAddressData(t *testing.T) {
	//proxy added later
	balance := GetBalance("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(balance.Balance)
<<<<<<< HEAD
	log.Println(DogeValueStrToInt(balance.Balance))
	received := GetReceived("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(received.Received)
	log.Println(DogeValueStrToInt(received.Received))
	sent := GetSent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(sent.Sent)
	log.Println(DogeValueStrToInt(sent.Sent))
=======
	received := GetReceived("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(received.Received)
	sent := GetSent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(sent.Sent)
>>>>>>> 0a15b0c1e9a828faf9fd8e5dbbf76b5fca9f34df
	unspent := GetUnspent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
	log.Println(unspent.UnspentOutputs)
	for _, row := range unspent.UnspentOutputs {
		log.Printf("%v %v", row.TxHash, row.TxOutputN)
	}
<<<<<<< HEAD
=======
	//	for _, row := range dat {
	//		log.Printf("%v %v", row.TxHash, row.TxOutputN)
	//	}
>>>>>>> 0a15b0c1e9a828faf9fd8e5dbbf76b5fca9f34df
}
