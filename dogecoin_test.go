//update v0.9
package dogecoin

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/alivanz/go-crypto/bitcoin"
)

func TestGetUnspentData(t *testing.T) {
	privkey := "3cd0560f5b27591916c643a0b7aa69d03839380a738d2e912990dcc573715d2c"
	privkeybin, _ := hex.DecodeString(privkey)
	wallet, _ := bitcoin.NewWallet(privkeybin)
	myaddress := WalletToAddress(wallet)
	destaddress := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	sendvalue := uint64(5000000000)
	balance := GetBalance(myaddress)
	log.Printf("myaddress    : %v\n", myaddress)
	log.Printf("balance      : %v\n", balance.Balance)
	log.Printf("destination  : %v\n", destaddress)
	log.Printf("sendvalue    : %v\n", IntToStr(sendvalue))
	if CanSpent(balance, sendvalue) != "saldo tidak mencukupi" {
		unspent := GetUnspent(myaddress)
		OrderUnspent(&unspent)
		dest := make([]Destination, 0)
		dest = append(dest, Destination{destaddress, sendvalue})
		numindex := ChangeUnspent(sendvalue, &unspent, &dest)
		signtx := CreateSignedTransaction(unspent, dest, wallet, numindex)
		log.Printf("signtxhex    : %v\n", signtx)
		bin, _ := hex.DecodeString(signtx)
		if err := DogeBroadcaster.Broadcast(bin); err != nil {
			log.Print(err)
			t.Fail()
			return
		}
	} else {
		log.Println(CanSpent(balance, sendvalue))
	}
}
