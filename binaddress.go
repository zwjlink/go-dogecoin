package dogecoin

func BinAddressNetworkCode(address []byte) byte {
	return address[0]
}

func BinAddressCheckSum(address []byte) []byte {
	return address[21:25]
}

func BinAddressPubKeyHash(address []byte) []byte {
	return address[1:21]
}
