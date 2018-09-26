package dogecoin

import "github.com/freddyisman/go-dogecoin/base58"

func AddressToPubKeyHash(address string) ([]byte, error) {
	return base58.Decode(address)
}
