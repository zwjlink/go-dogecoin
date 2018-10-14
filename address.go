//update v0.9
package dogecoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	crypto "github.com/alivanz/go-crypto"
	"github.com/alivanz/go-crypto/bitcoin/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	OP_DUP         = "76"
	OP_HASH160     = "a9"
	OP_EQUALVERIFY = "88"
	OP_CHECKSIG    = "ac"
	OP_TRUE        = "51"
	addrID         = "1e"
)

func Compressed(x, y string, compress int) string {
	var pubkey string
	if compress == 0 {
		pubkey = "04" + x + y
	} else if compress == 1 {
		suffix, err := strconv.ParseUint(y[len(y)-1:len(y)], 16, 64)
		ErrorCheck(err)
		if suffix%2 == 0 {
			pubkey = "02" + x
		} else {
			pubkey = "03" + x
		}
	}
	return pubkey
}

func WalletToAddress(wallet crypto.Wallet) string {
	var binaddress, pubkey bytes.Buffer
	wpubkey, _ := wallet.PubKey()
	x := fmt.Sprintf("%x", wpubkey.X)
	y := fmt.Sprintf("%x", wpubkey.Y)
	pubkey.WriteString(Compressed(x, y, 1))
	ripemd := ripemd160.New()
	pubkeybyte, err := hex.DecodeString(pubkey.String())
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
