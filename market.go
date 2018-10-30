package dogecoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Market struct {
	Data struct {
		Quotes struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quotes"`
	} `json:"data"`
}

func USDBasedFee(coinID string, USDfee float64) uint64 {
	resp, err := http.Get(fmt.Sprintf("https://api.coinmarketcap.com/v2/ticker/%v/?convert=USD", coinID))
	ErrorCheck(err)
	data, err := ioutil.ReadAll(resp.Body)
	ErrorCheck(err)
	var market Market
	err = json.Unmarshal(data, &market)
	ErrorCheck(err)
	coinfee := USDfee / market.Data.Quotes.USD.Price
	return uint64(coinfee)
}
