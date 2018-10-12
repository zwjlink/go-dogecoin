package dogecoin

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/alivanz/go-crypto/bitcoin"
)

type broadcastdoge struct{}

var DogeBroadcaster bitcoin.Broadcaster = broadcastdoge{}

func (broadcastdoge) Broadcast(signedtx []byte) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	data := make(map[string]string)
	data["tx"] = hex.EncodeToString(signedtx)
	bin, _ := json.Marshal(data)
	// data.Add("coin_symbol", "doge")
	// data.Add("csrfmiddlewaretoken", token)
	request, err := http.NewRequest("POST", "https://api.blockcypher.com/v1/doge/main/txs/push", bytes.NewBuffer(bin))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "text/json")
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode == 400 {
		var ret map[string]string
		msg, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(msg, &ret)
		return fmt.Errorf("%s", ret["error"])
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("Status code %d", resp.StatusCode)
	}
	return nil
}
