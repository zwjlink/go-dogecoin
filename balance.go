package dogecoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DogeChainBalance struct {
	Balance string `json:"balance"`
	Success int    `json:"success"`
}

func (doge Doge) GetBalance(address string) uint64 {
	resp, err := http.Get(fmt.Sprintf("https://dogechain.info/api/v1/address/balance/%v", address))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var balance DogeChainBalance
	err = json.Unmarshal(data, &balance)
	ErrorCheck(err)
	return StrToInt(balance.Balance)
}
