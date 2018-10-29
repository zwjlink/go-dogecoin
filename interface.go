package dogecoin

type Currency interface {
	CreateCoin(pubkeyhash string) Coin
	Broadcast(signtx string) error
}

type Doge struct{}
type Dash struct{}
