package utils

import "strconv"

func ToInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

func ToIntDefault(s string, _default int) int {
	num, err := strconv.Atoi(s)
	if err == nil {
		return num
	}
	return _default
}
