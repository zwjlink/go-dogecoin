package dogecoin

import (
	"fmt"
	"strconv"
	"strings"
)

func DogeValueStrToInt(doge_value string) uint64 {
	doge_string_value := fmt.Sprint(strings.Replace(doge_value, ".", "", -1))
	value, err := strconv.ParseUint(doge_string_value, 10, 64)
	ErrorCheck(err)
	return value
}
