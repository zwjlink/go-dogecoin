package dogecoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DogeChainUnspent struct {
	UnspentOutputs []struct {
		TxHash        string `json:"tx_hash"`
		TxOutputN     int    `json:"tx_output_n"`
		Script        string `json:"script"`
		Value         string `json:"value"`
		Confirmations int    `json:"confirmations"`
		Address       string `json:"address"`
	} `json:"unspent_outputs"`
	Success int `json:"success"`
}

type Unspent struct {
	TxHash    string
	TxOutputN int
	Script    string
	Value     uint64
}

func (doge Doge) GetUnspent(address string) []Unspent {
	resp, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/unspent/%v", address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var dogeunspent DogeChainUnspent
	err = json.Unmarshal(data, &dogeunspent)
	ErrorCheck(err)
	unspent := make([]Unspent, 0)
	for i := 0; i < len(dogeunspent.UnspentOutputs); i++ {
		unspent = append(unspent, Unspent{
			dogeunspent.UnspentOutputs[i].TxHash,
			dogeunspent.UnspentOutputs[i].TxOutputN,
			dogeunspent.UnspentOutputs[i].Script,
			StrToInt(dogeunspent.UnspentOutputs[i].Value),
		})
	}
	OrderUnspent(&unspent)
	return unspent
}
