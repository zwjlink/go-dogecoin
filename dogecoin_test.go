package dogecoin

import (
	"log"
	"testing"
)

func TestGetUnspentData(t *testing.T) {
	unspent := GetUnspent("DMr3fEiVrPWFpoCWS958zNtqgnFb7QWn9D")
	OrderUnspent(&unspent)
	input, change := InputTemplate(1500000000000, unspent)
	for _, row := range unspent.UnspentOutputs {
		log.Printf("%v %v %v", row.TxHash, row.TxOutputN, row.Value)
	}
	log.Print(input)
	log.Print(change)
}

// unspent := GetUnspent("DMr3fEiVrPWFpoCWS958zNtqgnFb7QWn9D")
// for _, row := range unspent.UnspentOutputs {
// 	log.Printf("%v %v %v", row.TxHash, row.TxOutputN, row.Value)
// }
// OrderUnspent(&unspent)
// log.Println("")

// reversed := ReverseHex("76a91499b1ebcfc11a13df5161aba8160460fe1601d54188ac")
// log.Println(reversed)

// bytepresent, err := hex.DecodeString("ffffffff")
// ErrorCheck(err)
// log.Printf("%v", bytepresent)

// BinAddress, err := base58.Decode("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// ErrorCheck(err)
// p2pkh := P2PKH(BinAddressPubKeyHash(BinAddress))
// log.Println(hex.EncodeToString(p2pkh))
// unspent := GetUnspent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Printf("%v", unspent.UnspentOutputs[0].Script)

// BinAddress, err := base58.Decode("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// ErrorCheck(err)
// log.Println(hex.EncodeToString(BinAddress)))
// log.Println(hex.EncodeToString(BinAddressNetworkCode(BinAddress)))
// log.Println(hex.EncodeToString(BinAddressPubKeyHash(BinAddress)))
// log.Println(hex.EncodeToString(BinAddressCheckSum(BinAddress)))
// log.Println(hex.EncodeToString(Hash(BinAddress[0:21])))
// unspent := GetUnspent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Printf("%v %v", unspent.UnspentOutputs[0].Script, unspent.UnspentOutputs[0].Address)

//proxy added later
// balance := GetBalance("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Println(balance.Balance)
// log.Println(DogeValueStrToInt(balance.Balance))
// received := GetReceived("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Println(received.Received)
// log.Println(DogeValueStrToInt(received.Received))
// sent := GetSent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Println(sent.Sent)
// log.Println(DogeValueStrToInt(sent.Sent))
// unspent := GetUnspent("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Println(unspent.UnspentOutputs)
// for _, row := range unspent.UnspentOutputs {
// 	log.Printf("%v %v", row.TxHash, row.TxOutputN)
// }
// ScriptPubKey, err := AddressToScriptPubKey("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// log.Println(hex.EncodeToString(Hash(ScriptPubKey[0:21])))
// log.Println(hex.EncodeToString(ScriptPubKey[21:25]))
// ErrorCheck(err)
// log.Println(hex.EncodeToString(ScriptPubKey))
// address, err := bitcoin.AddressParseBase58("DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi")
// ErrorCheck(err)
// log.Println(address.Version())
// log.Println(hex.EncodeToString(address.PubKeyHash()))
