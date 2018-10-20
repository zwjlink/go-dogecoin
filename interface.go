package dogecoin

type Currency interface {
	GetBalance(address string) uint64
	GetUnspent(address string) []Unspent
	Broadcast(signtx string) error
	CreateCoin(pubkeyhash string) Coin
}

type Doge struct{}
