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

func (doge Doge) Broadcast(signtx string) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	// menerima data hex transaksi yang sudah di signature dalam bentuk string
	data := make(map[string]string)
	data["tx"] = signtx
	// konversi data string ke format json
	bin, _ := json.Marshal(data)
	// push data yang sudah dikonversi ke link tujuan, sesuai format API dari blockcypher
	request, err := http.NewRequest("POST", "https://api.blockcypher.com/v1/doge/main/txs/push", bytes.NewBuffer(bin))
	ErrorCheck(err)
	request.Header.Add("Content-Type", "text/json")
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
