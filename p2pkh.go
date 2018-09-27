package dogecoin

const (
	OP_DUP         byte = 0x76
	OP_HASH160     byte = 0xa9
	OP_EQUALVERIFY byte = 0x88
	OP_CHECKSIG    byte = 0xac
	//OP_TRUE        byte = 0x51
)

func P2PKH(pubkeyhash []byte) []byte {
	p2pkh := make([]byte, 25)
	p2pkh[0] = OP_DUP
	p2pkh[1] = OP_HASH160
	p2pkh[2] = byte(len(pubkeyhash))
	for i := 0; i < 20; i++ {
		p2pkh[3+i] = pubkeyhash[i]
	}
	p2pkh[23] = OP_EQUALVERIFY
	p2pkh[24] = OP_CHECKSIG
	return p2pkh
}

// Anyone-Can-Spend Outputs
// func AnyoneCanSpent() []byte {
// 	return []byte{OP_TRUE}
// }
