package dogecoin

type Coin struct {
	Version string
	Fee     uint64
	Address string
	Balance uint64
	Unspent []Unspent
}

func (doge Doge) CreateCoin(pubkeyhash string) Coin {
	var dogecoin Coin
	dogecoin.Version = "01000000"
	// initial dogefee per approximately 150 character or one output is 0.25 doge, input value of fee in USD satoshi
	dogecoin.Fee = USDBasedFee("74", float64(100000))
	dogecoin.Address = GetAddress(pubkeyhash, "1e")
	dogecoin.Balance, dogecoin.Unspent = GetBlockCypherChain("doge", dogecoin.Address)
	return dogecoin
}

func (doge Doge) Broadcast(signtx string) {
	ErrorCheck(Broadcasting("doge", signtx))
}

func (dash Dash) CreateCoin(pubkeyhash string) Coin {
	var dashcoin Coin
	dashcoin.Version = "02000000"
	dashcoin.Fee = USDBasedFee("131", float64(3000000))
	dashcoin.Address = GetAddress(pubkeyhash, "4c")
	dashcoin.Balance, dashcoin.Unspent = GetBlockCypherChain("dash", dashcoin.Address)
	return dashcoin
}

func (dash Dash) Broadcast(signtx string) {
	ErrorCheck(Broadcasting("dash", signtx))
}
