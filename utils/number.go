package utils

import (
	"strconv"

	"github.com/miladbonakdar/tp-rate-review/fail"
)

func ParseToUint8(s string) (uint8, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, fail.NewFailByError(400, err, "ParseToUint8")
	}
	return uint8(val), nil
}
