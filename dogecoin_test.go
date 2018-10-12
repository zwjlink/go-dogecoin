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
	privkey := "e105500a65cd0eda7ec6784a27a09f20c725ade74ec7d1bd96d09318d0ed43a4"
	sendvalue := uint64(50000000)
	destaddress := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	privkeybin, _ := hex.DecodeString(privkey)
	wallet, _ := bitcoin.NewWallet(privkeybin)
	wpubkey, _ := wallet.PubKey()
	x := fmt.Sprintf("%x", wpubkey.X)
	y := fmt.Sprintf("%x", wpubkey.Y)
	pubkey.WriteString(Compressed(x, y, 1))
	Address := PubKeyToAddress(pubkey.String())
	balance := GetBalance(Address)
	log.Printf("NetworkID    : %v\n", addrID)
	log.Printf("mypublickey  : %v\n", pubkey.String())
	log.Printf("myaddress    : %v\n", Address)
	log.Printf("destination  : %v\n", destaddress)
	log.Printf("sendvalue    : %v\n", IntToStr(sendvalue))
	log.Printf("balance      : %v\n", balance.Balance)
	unspent := GetUnspent(Address)
	OrderUnspent(&unspent)
	dest := make([]Destination, 0)
	dest = append(dest, Destination{destaddress, sendvalue})
	rawtx, _ := CreateRawTransaction(unspent, balance, &dest)
	log.Printf("rawtxhex     : %v\n", rawtx)
	if rawtx != "saldo tidak mencukupi" {
		rawtxbyte, _ := hex.DecodeString(rawtx)
		rawtxhash := Hash(rawtxbyte)
		r, s, _ := wallet.Sign(rawtxhash)
		r_correct := SignCorrect(fmt.Sprintf("%x", r))
		s_correct := SignCorrect(fmt.Sprintf("%x", s))
		scriptsig := ScriptSig(r_correct, s_correct, pubkey.String())
		signtx, _ := CreateSignedTransaction(unspent, balance, dest, scriptsig)
		valid := ecdsa.Verify(&wpubkey, rawtxhash, r, s)
		log.Printf("rawtxhash    : %v", hex.EncodeToString(rawtxhash))
		log.Printf("signtxhex    : %v\n", signtx)
		log.Println("signature verified:", valid)
		bin, _ := hex.DecodeString(signtx)
		if err := DogeBroadcaster.Broadcast(bin); err != nil {
			log.Print(err)
			t.Fail()
			return
		}
	}
}
