package dogecoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Broadcasting(coin string, signtx string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// menerima data hex transaksi yang sudah di signature dalam bentuk string
	data := make(map[string]string)
	data["tx"] = signtx
	// konversi data string ke format json
	bin, _ := json.Marshal(data)
	// request push ke link tujuan, sesuai dokumentasi API dari blockcypher
	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.blockcypher.com/v1/%v/main/txs/push", coin), bytes.NewBuffer(bin))
	ErrorCheck(err)
	// just header, not necessary
	request.Header.Add("Content-Type", "text/json")
	// send the request to the link and return the link response
	resp, err := client.Do(request)
	ErrorCheck(err)
	// respon jika gagal
	if resp.StatusCode == 400 {
		var ret map[string]string
		msg, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(msg, &ret)
		return fmt.Errorf("%s", ret["error"])
	}
	// respon jika berhasil
	if resp.StatusCode != 200 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("Status code %d", resp.StatusCode)
	}
	return nil
}
