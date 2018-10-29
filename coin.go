package dogecoin

//in USD satoshi
const USDfeeDoge = 100000

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
	//find a link to convert USD fee to Doge fee
	dogecoin.fee = 25000000
	dogecoin.address = GetAddress(pubkeyhash, dogecoin.ID)
	dogecoin.balance, dogecoin.unspent = GetBlockCypherChain("doge", dogecoin.address)
	return dogecoin
}

func (doge Doge) Broadcast(signtx string) {
	ErrorCheck(Broadcasting("doge", signtx))
}
