package utils

import "strconv"

func ParseToUint8(s string) (uint8, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint8(val), nil
}
