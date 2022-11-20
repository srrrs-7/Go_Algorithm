package utils

import "strconv"

func ArrayToInteger() {

}

func arrayToInteger() {
	st := ""
	for _, i := range results {
		s := strconv.Itoa(i)
		st += s
	}

	res, err := strconv.Atoi(st)
	if err != nil {
		return -1
	}
}