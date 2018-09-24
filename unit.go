package dogecoin

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Unit uint64

const Doge = Unit(100000000)

func (u Unit) String() string {
	return fmt.Sprintf("%d.%d", u/Doge, u%Doge)
}
func (u *Unit) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	ss := strings.Split(s, ".")
	if len(ss) == 1 {
		x, err := strconv.ParseUint(ss[0], 10, 64)
		if err != nil {
			return err
		}
		*u = Unit(x * uint64(Doge))
		return nil
	} else if len(ss) == 2 {
		x, err := strconv.ParseUint(ss[0], 10, 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseUint(ss[1], 10, 64)
		if err != nil {
			return err
		}
		*u = Unit(x*uint64(Doge) + y)
		return nil
	}
	return fmt.Errorf("titik bnyk")
}
