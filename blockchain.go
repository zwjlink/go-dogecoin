package dogecoin

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

type DogeChainBalance struct {
	Balance string `json:"balance"`
}

type DogeChainUnspent struct {
	UnspentOutputs []struct {
		TxHash    string `json:"tx_hash"`
		TxOutputN int    `json:"tx_output_n"`
		Script    string `json:"script"`
		Value     string `json:"value"`
	} `json:"unspent_outputs"`
}

type Unspent struct {
	TxHash    string
	TxOutputN int
	Script    string
	Value     uint64
}

func GetDogeChain(address string) (uint64, []Unspent) {
	resp_balance, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/address/balance/%v", address))
	ErrorCheck(err)
	resp_unspent, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/unspent/%v", address))
	ErrorCheck(err)
	data_balance, err := ioutil.ReadAll(resp_balance.Body)
	ErrorCheck(err)
	data_unspent, err := ioutil.ReadAll(resp_unspent.Body)
	ErrorCheck(err)
	var doge_balance DogeChainBalance
	var doge_unspent DogeChainUnspent
	ErrorCheck(json.Unmarshal(data_balance, &doge_balance))
	ErrorCheck(json.Unmarshal(data_unspent, &doge_unspent))
	unspent := make([]Unspent, 0)
	for i := 0; i < len(doge_unspent.UnspentOutputs); i++ {
		unspent = append(unspent, Unspent{
			doge_unspent.UnspentOutputs[i].TxHash,
			doge_unspent.UnspentOutputs[i].TxOutputN,
			doge_unspent.UnspentOutputs[i].Script,
			StrToInt(doge_unspent.UnspentOutputs[i].Value),
		})
	}
	OrderUnspent(&unspent)
	return StrToInt(doge_balance.Balance), unspent
}

func GetBlockCypherChain(coin string, address string) (uint64, []Unspent) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://api.blockcypher.com/v1/%v/main/addrs/%v?unspentOnly=true", coin, address), nil)
	resp, err := client.Do(req)
	// if no response for a given time, then go to dogechain
	if err != nil {
		return GetDogeChain(address)
	}
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
