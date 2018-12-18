package dogecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	crypto "github.com/alivanz/go-crypto"
)

const (
	locktime = "00000000"
	sighash  = "01000000"
)

// membuat signature
func CreateSignature(coindata Coin, dest []Destination, wallet crypto.Wallet, numindex int, posindex int) string {
	var rawtx, pubkey bytes.Buffer
	// membuat hex input untuk hex transaksi raw
	inputstr := InputTemplate(coindata, nil, nil, numindex, posindex)
	// membuat hex output untuk hex transaksi raw
	outputstr := OutputTemplate(dest)
	// pembentukan hex transaksi raw
	rawtx.WriteString(coindata.Version)
	rawtx.WriteString(inputstr)
	rawtx.WriteString(outputstr)
	rawtx.WriteString(locktime)
	rawtx.WriteString(sighash)
	// konversi hex transaksi raw ke dalam tipe byte untuk keperluan hashing
	rawtxbyte, _ := hex.DecodeString(rawtx.String())
	// melakukan hashing pada hex transaksi raw sehingga diperoleh hash transaksi raw
	rawtxhash := Hash(rawtxbyte)
	// mengambil pubkey yang belum dikompresi dari object wallet
	wpubkey, _ := wallet.PubKey()
	// signature dilakukan pada hash transaksi raw, sehingga diperoleh komponen signature r dan s
	r, s, _ := wallet.Sign(rawtxhash)
	// koreksi komponen r dan s untuk mencegah nilai negatif
	r_correct := SignCorrect(fmt.Sprintf("%x", r))
	s_correct := SignCorrect(fmt.Sprintf("%x", s))
	// ekstraksi nilai x dan y dari pubkey yang masih belum dikompresi
	x := fmt.Sprintf("%x", wpubkey.X)
	y := fmt.Sprintf("%x", wpubkey.Y)
	// membentuk pubkey versi kompresi-nya
	pubkey.WriteString(Compressed(x, y, 1))
	// pembentukan final signature dari komponen r, s dan pubkey yang telah dikompresi
	scriptsig := ScriptSig(r_correct, s_correct, pubkey.String())
	return scriptsig
}

func CreateSignedTransaction(coindata Coin, dest []Destination, wallet crypto.Wallet, numindex int) string {
	var signedtx bytes.Buffer
	// menbentuk hex input untuk hex transaksi yang di signature
	inputstr := InputTemplate(coindata, dest, wallet, numindex, -1)
	// menbentuk hex output untuk hex transaksi yang di signature
	outputstr := OutputTemplate(dest)
	// pembentukan hex transaksi yang sudah di signature
	signedtx.WriteString(coindata.Version)
	signedtx.WriteString(inputstr)
	signedtx.WriteString(outputstr)
	signedtx.WriteString(locktime)
	return signedtx.String()
}
