package dogecoin

import (
	"github.com/alivanz/go-crypto/bitcoin/base58"
)

func AddressToPubKeyHash(address string) ([]byte, error) {
	return base58.Decode(address)
}
