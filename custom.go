//update v0.9
package dogecoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const signlim = 0x7f
const satoshi = 100000000

func SignCorrect(sign string) string {
	for len(sign) < 64 {
		sign = "0" + sign
	}
	prefix, err := strconv.ParseUint(sign[:2], 16, 64)
	ErrorCheck(err)
	if prefix >= signlim {
		sign = "00" + sign
	}
	return sign
}

func EvenCorrect(num int) string {
	var numstring bytes.Buffer
	if (len(fmt.Sprintf("%x", num)) % 2) > 0 {
		numstring.WriteString("0" + fmt.Sprintf("%x", num))
	} else {
		numstring.WriteString(fmt.Sprintf("%x", num))
	}
	return numstring.String()
}

func VarInt(num int) string {
	var numstring, numfinal bytes.Buffer
	numstring.WriteString(EvenCorrect(num))
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

func IntToStr(doge_value uint64) string {
	return fmt.Sprintf("%v.%v", doge_value/satoshi, doge_value%satoshi)
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

func Hash(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:]
}

func ErrorCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}
