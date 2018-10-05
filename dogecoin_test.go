package dogecoin

import (
	"log"
	"testing"
)

func TestGetUnspentData(t *testing.T) {
	unspent := GetUnspent("DMr3fEiVrPWFpoCWS958zNtqgnFb7QWn9D")
	OrderUnspent(&unspent)
	dest := make([]Destination, 0)
	dest = append(dest, Destination{"DGZvtQkZo8dGhpn8DqAHNUjmQVrbAFGHQi", 257000000000})
	dest = append(dest, Destination{"DRsLzWNubViGtLoxXnaNaMbvTiadMMsUYL", 125000000000})
	rawtransaction := CreateRawTransaction(unspent, &dest)
	log.Println(rawtransaction)
}
