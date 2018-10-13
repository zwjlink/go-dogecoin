//update v0.9
package dogecoin

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/alivanz/go-crypto/bitcoin"
)

func TestGetUnspentData(t *testing.T) {
	privkey := "e105500a65cd0eda7ec6784a27a09f20c725ade74ec7d1bd96d09318d0ed43a4"
	myaddress := "DGaL1Bm3YmWCnuz4j3BSPQZSNYGysgX9ZL"
	privkeybin, _ := hex.DecodeString(privkey)
	wallet, _ := bitcoin.NewWallet(privkeybin)
	destaddress := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	sendvalue := uint64(3000000000)
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
