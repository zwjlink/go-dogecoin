//update v0.9
package dogecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	crypto "github.com/alivanz/go-crypto"
)

const (
	version  = "01000000"
	locktime = "00000000"
	sighash  = "01000000"
)

func CreateSignature(unspent DogechainUnspent, dest []Destination, wallet crypto.Wallet, numindex int, posindex int) string {
	var rawtx, pubkey bytes.Buffer
	inputstr := InputTemplate(unspent, nil, nil, numindex, posindex)
	outputstr := OutputTemplate(dest)
	rawtx.WriteString(version)
	rawtx.WriteString(inputstr)
	rawtx.WriteString(outputstr)
	rawtx.WriteString(locktime)
	rawtx.WriteString(sighash)
	rawtxbyte, _ := hex.DecodeString(rawtx.String())
	rawtxhash := Hash(rawtxbyte)
	r, s, _ := wallet.Sign(rawtxhash)
	wpubkey, _ := wallet.PubKey()
	r_correct := SignCorrect(fmt.Sprintf("%x", r))
	s_correct := SignCorrect(fmt.Sprintf("%x", s))
	x := fmt.Sprintf("%x", wpubkey.X)
	y := fmt.Sprintf("%x", wpubkey.Y)
	pubkey.WriteString(Compressed(x, y, 1))
	scriptsig := ScriptSig(r_correct, s_correct, pubkey.String())
	return scriptsig
}

func CreateSignedTransaction(unspent DogechainUnspent, dest []Destination, wallet crypto.Wallet, numindex int) string {
	var signedtx bytes.Buffer
	inputstr := InputTemplate(unspent, dest, wallet, numindex, -1)
	outputstr := OutputTemplate(dest)
	signedtx.WriteString(version)
	signedtx.WriteString(inputstr)
	signedtx.WriteString(outputstr)
	signedtx.WriteString(locktime)
	return signedtx.String()
}
