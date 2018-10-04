package dogecoin

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(doge_value string) uint64 {
	var value uint64
	var err error
	doge_string_value := fmt.Sprint(strings.Replace(doge_value, ".", "", -1))
	if doge_string_value == "" {
		value = 0
	} else {
		value, err = strconv.ParseUint(doge_string_value, 10, 64)
		ErrorCheck(err)
	}
	return value
}

func ReverseHex(hexa string) string {
	bytes, err := hex.DecodeString(hexa)
	ErrorCheck(err)
	var temp byte
	for i := 0; i < (len(bytes) - (i + 1)); i++ {
		temp = bytes[i]
		bytes[i] = bytes[len(bytes)-(i+1)]
		bytes[len(bytes)-(i+1)] = temp
	}
	reversed := hex.EncodeToString(bytes)
	return reversed
}
