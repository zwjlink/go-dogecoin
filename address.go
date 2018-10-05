//update v0.9
package dogecoin

import (
	"bytes"
	"fmt"
)

const (
	OP_DUP         = "76"
	OP_HASH160     = "a9"
	OP_EQUALVERIFY = "88"
	OP_CHECKSIG    = "ac"
	OP_TRUE        = "51"
)

func BinAddressNetworkID(binaddress string) string {
	return binaddress[0:1]
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
