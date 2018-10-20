package dogecoin

type Coin struct {
	ID      string
	version string
	fee     uint64
	address string
	balance uint64
	unspent []Unspent
}

func (doge Doge) CreateCoin(pubkeyhash string) Coin {
	var dogecoin Coin
	dogecoin.ID = "1e"
	dogecoin.version = "01000000"
	dogecoin.fee = 25000000
	dogecoin.address = GetAddress(pubkeyhash, dogecoin.ID)
	dogecoin.balance = doge.GetBalance(dogecoin.address)
	dogecoin.unspent = doge.GetUnspent(dogecoin.address)
	return dogecoin
}
