package dogecoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DogechainBalance struct {
	Balance string `json:"balance"`
	Success int    `json:"success"`
}

type DogechainReceived struct {
	Received string `json:"received"`
	Success  int    `json:"success"`
}

type DogechainSent struct {
	Sent    string `json:"sent"`
	Success int    `json:"success"`
}

type DogechainUnspent struct {
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

func GetBalance(address string) DogechainBalance {
	resp, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/address/balance/%v", address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var balance DogechainBalance
	err = json.Unmarshal(data, &balance)
	ErrorCheck(err)
	return balance
}

func GetReceived(address string) DogechainReceived {
	resp, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/address/received/%v", address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var received DogechainReceived
	err = json.Unmarshal(data, &received)
	ErrorCheck(err)
	return received
}

func GetSent(address string) DogechainSent {
	resp, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/address/sent/%v", address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var sent DogechainSent
	err = json.Unmarshal(data, &sent)
	ErrorCheck(err)
	return sent
}

func GetUnspent(address string) DogechainUnspent {
	resp, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/unspent/%v", address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var unspent DogechainUnspent
	err = json.Unmarshal(data, &unspent)
	ErrorCheck(err)
	return unspent
}
