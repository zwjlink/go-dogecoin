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

// Anyone-Can-Spend Outputs
// func AnyoneCanSpent() []byte {
// 	return []byte{OP_TRUE}
// }
