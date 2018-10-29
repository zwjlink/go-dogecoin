package dogecoin

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/btcsuite/btcutil/base58"
)

type BlockCypherChain struct {
	Balance uint64 `json:"balance"`
	Txrefs  []struct {
		TxHash    string `json:"tx_hash"`
		TxOutputN int    `json:"tx_output_n"`
		Value     uint64 `json:"value"`
	} `json:"txrefs"`
}

type Unspent struct {
	TxHash    string
	TxOutputN int
	Script    string
	Value     uint64
}

func GetBlockCypherChain(coin string, address string) (uint64, []Unspent) {
	resp, err := http.Get(fmt.Sprintf("http://api.blockcypher.com/v1/%v/main/addrs/%v?unspentOnly=true", coin, address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var blockchain BlockCypherChain
	err = json.Unmarshal(data, &blockchain)
	ErrorCheck(err)
	binaddress := base58.Decode(address)
	script := P2PKH(hex.EncodeToString(binaddress[1:21]))
	unspent := make([]Unspent, 0)
	for i := 0; i < len(blockchain.Txrefs); i++ {
		unspent = append(unspent, Unspent{
			blockchain.Txrefs[i].TxHash,
			blockchain.Txrefs[i].TxOutputN,
			script,
			blockchain.Txrefs[i].Value,
		})
	}
	OrderUnspent(&unspent)
	return blockchain.Balance, unspent
}
