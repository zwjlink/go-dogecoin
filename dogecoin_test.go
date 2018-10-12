//update v0.9
package dogecoin

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/alivanz/go-crypto/bitcoin"
)

func TestGetUnspentData(t *testing.T) {
	var pubkey bytes.Buffer
	var balance DogechainBalance
	privkey := "e105500a65cd0eda7ec6784a27a09f20c725ade74ec7d1bd96d09318d0ed43a4"
	privkeybin, _ := hex.DecodeString(privkey)
	wallet, err := bitcoin.NewWallet(privkeybin)
	ErrorCheck(err)
	wpubkey, _ := wallet.PubKey()
	x := fmt.Sprintf("%x", wpubkey.X)
	y := fmt.Sprintf("%x", wpubkey.Y)
	pubkey.WriteString(Compressed(x, y, 1))
	Address := PubKeyToAddress(pubkey.String())
	sendvalue := uint64(1300000000)
	destaddress := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	log.Printf("NetworkID    : %v\n", addrID)
	log.Printf("mypublickey  : %v\n", pubkey.String())
	log.Printf("myaddress    : %v\n", Address)
	log.Printf("destination  : %v\n", destaddress)
	log.Printf("sendvalue    : %v\n", sendvalue)
	balance = GetBalance(Address)
	log.Printf("balance      : %v\n", balance.Balance)
	unspent := GetUnspent(Address)
	OrderUnspent(&unspent)
	dest := make([]Destination, 0)
	dest = append(dest, Destination{destaddress, sendvalue})
	rawtx, change := CreateRawTransaction(unspent, balance, &dest)
	log.Printf("rawtxhex     : %v\n", rawtx)
	if rawtx != "saldo tidak mencukupi" {
		log.Printf("change       : %v\n", change)
		rawtxbyte, err := hex.DecodeString(rawtx)
		ErrorCheck(err)
		rawtxhash := Hash(rawtxbyte)
		log.Printf("rawtxhash    : %v", hex.EncodeToString(rawtxhash))
		r, s, err := wallet.Sign(rawtxhash)
		ErrorCheck(err)
		r_correct := SignCorrect(fmt.Sprintf("%x", r))
		s_correct := SignCorrect(fmt.Sprintf("%x", s))
		scriptsig := ScriptSig(r_correct, s_correct, pubkey.String())
		signtx, change := CreateSignedTransaction(unspent, balance, dest, scriptsig)
		log.Printf("signtxhex    : %v\n", signtx)
		log.Printf("change       : %v\n", change)
		valid := ecdsa.Verify(&wpubkey, rawtxhash, r, s)
		log.Println("signature verified:", valid)
		bin, _ := hex.DecodeString(signtx)
		if err := DogeBroadcaster.Broadcast(bin); err != nil {
			log.Print(err)
			t.Fail()
			return
		}
	}
}
