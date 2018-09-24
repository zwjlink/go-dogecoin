package dogecoin

import "log"

func ErrorCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}
