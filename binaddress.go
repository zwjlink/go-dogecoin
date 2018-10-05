package dogecoin

func BinAddressNetworkID(binaddress string) string {
	return binaddress[0:1]
}

func BinAddressCheckSum(binaddress string) string {
	return binaddress[42:50]
}

func BinAddressPubKeyHash(binaddress string) string {
	return binaddress[2:42]
}
