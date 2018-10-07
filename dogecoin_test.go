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

	"github.com/freddyisman/go-dogecoin/base58"
)

func TestGetUnspentData(t *testing.T) {
	var pubkey bytes.Buffer
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubkey.WriteString("04")
	pubkey.WriteString(fmt.Sprintf("%x%x", (*privkey).PublicKey.X, (*privkey).PublicKey.Y))
	exbinaddr, err := base58.Decode("DBHdUWoMQAAt4CgCfWztwRovxhQQ9qo5Um")
	ErrorCheck(err)
	addrID := BinAddressNetworkID(hex.EncodeToString(exbinaddr))
	Address := PubKeyToAddress(pubkey.String(), addrID)
	log.Printf("myprivatekey : %x\n", privkey.D)
	log.Printf("mypublickey  : %v\n", pubkey.String())
	log.Printf("myaddress    : %v\n", Address)
	unspent := GetUnspent(Address)
	balance := GetBalance(Address)
	OrderUnspent(&unspent)
	dest := make([]Destination, 0)
	sendvalue := uint64(5700000000)
	destaddress := "DBHdUWoMQAAt4CgCfWztwRovxhQQ9qo5Um"
	dest = append(dest, Destination{destaddress, sendvalue})
	log.Printf("destination : %v\n", destaddress)
	log.Printf("sendvalue   : %v\n", sendvalue)
	rawtx, change := CreateRawTransaction(unspent, balance, &dest)
	log.Printf("rawtxhex  : %v\n", rawtx)
	if rawtx == "saldo tidak mencukupi" {
		log.Printf("balance   : %v\n", change)
	} else {
		rawtxbyte, err := hex.DecodeString(rawtx)
		ErrorCheck(err)
		rawtxhash := Hash(rawtxbyte)
		log.Printf("rawtxhash : %v", hex.EncodeToString(rawtxhash))
		r, s, err := ecdsa.Sign(rand.Reader, privkey, rawtxhash)
		ErrorCheck(err)
		scriptsig := ScriptSig(fmt.Sprintf("%x", r), fmt.Sprintf("%x", s), pubkey.String())
		signtx, change := CreateSignedTransaction(unspent, balance, dest, scriptsig)
		log.Printf("signtxhex : %v\n", signtx)
		log.Printf("change    : %v\n", change)
	}
}
