package utils

import "strconv"

func ToInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}
