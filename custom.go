//update v0.9
package dogecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func VarInt(num int) string {
	var numstring, numfinal bytes.Buffer
	if (len(fmt.Sprint(num)) % 2) > 0 {
		numstring.WriteString("0" + fmt.Sprintf("%x", num))
	} else {
		numstring.WriteString(fmt.Sprintf("%x", num))
	}
	switch {
	case num <= 0xfc:
		//do nothing
	case num <= 0xffff:
		numfinal.WriteString("fd")
	case num <= 0xffffffff:
		numfinal.WriteString("fe")
	default:
		numfinal.WriteString("ff")
	}
	numfinal.WriteString(ReverseHex(numstring.String()))
	return numfinal.String()
}

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
