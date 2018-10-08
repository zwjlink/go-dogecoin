//update v0.9
package dogecoin

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
	"time"
)

const (
	addrID       = "1e"
	uncompressed = "04"
)

func TestGetUnspentData(t *testing.T) {
	var pubkey bytes.Buffer
	var balance DogechainBalance
	random := rand.Reader
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), random)
	ErrorCheck(err)
	pubkey.WriteString(uncompressed)
	pubkey.WriteString(fmt.Sprintf("%x%x", (*privkey).PublicKey.X, (*privkey).PublicKey.Y))
	Address := PubKeyToAddress(pubkey.String(), addrID)
	sendvalue := uint64(270000000)
	destaddress := "DPAQVCUVQU1LKRkeKihjYb2gDiHoLteSwR"
	log.Printf("NetworkID    : %v\n", addrID)
	log.Printf("myprivatekey : %x\n", privkey.D)
	log.Printf("mypublickey  : %v\n", pubkey.String())
	log.Printf("myaddress    : %v\n", Address)
	log.Printf("destination  : %v\n", destaddress)
	log.Printf("sendvalue    : %v\n", sendvalue)
	duration := time.Duration(600) * time.Second
	time.Sleep(duration)
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
		r, s, err := ecdsa.Sign(random, privkey, rawtxhash)
		ErrorCheck(err)
		r_correct := SignCorrect(fmt.Sprintf("%x", r))
		s_correct := SignCorrect(fmt.Sprintf("%x", s))
		scriptsig := ScriptSig(r_correct, s_correct, pubkey.String())
		signtx, change := CreateSignedTransaction(unspent, balance, dest, scriptsig)
		log.Printf("signtxhex    : %v\n", signtx)
		log.Printf("change       : %v\n", change)
		valid := ecdsa.Verify(&privkey.PublicKey, rawtxhash, r, s)
		log.Println("signature verified:", valid)
	}
}
