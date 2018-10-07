//update v0.9
package dogecoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/freddyisman/go-dogecoin/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	OP_DUP         = "76"
	OP_HASH160     = "a9"
	OP_EQUALVERIFY = "88"
	OP_CHECKSIG    = "ac"
	OP_TRUE        = "51"
)

func PubKeyToAddress(pubkey string, addrID string) string {
	var binaddress bytes.Buffer
	ripemd := ripemd160.New()
	pubkeybyte, err := hex.DecodeString(pubkey)
	ErrorCheck(err)
	firsthash := sha256.Sum256(pubkeybyte)
	ripemd.Write(firsthash[:])
	pubkeyhash := ripemd.Sum(nil)
	binaddress.WriteString(addrID)
	binaddress.WriteString(hex.EncodeToString(pubkeyhash))
	binaddrnocek, err := hex.DecodeString(binaddress.String())
	ErrorCheck(err)
	checksum := Hash(binaddrnocek)[:4]
	binaddress.WriteString(hex.EncodeToString(checksum))
	binaddrbyte, err := hex.DecodeString(binaddress.String())
	ErrorCheck(err)
	address := base58.Encode(binaddrbyte)
	return address
}

func BinAddressNetworkID(binaddress string) string {
	return binaddress[:2]
}

func BinAddressCheckSum(binaddress string) string {
	return binaddress[42:50]
}

func BinAddressPubKeyHash(binaddress string) string {
	return binaddress[2:42]
}

func P2PKH(pubkeyhash string) string {
	var p2pkh bytes.Buffer
	p2pkh.WriteString(OP_DUP)
	p2pkh.WriteString(OP_HASH160)
	p2pkh.WriteString(fmt.Sprintf("%x", len(pubkeyhash)/2))
	p2pkh.WriteString(pubkeyhash)
	p2pkh.WriteString(OP_EQUALVERIFY)
	p2pkh.WriteString(OP_CHECKSIG)
	return p2pkh.String()
}
