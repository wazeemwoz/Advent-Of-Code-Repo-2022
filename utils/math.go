package utils

func Max(ints ...int) int {
	max := ints[0]
	for _, v := range ints {
		if max < v {
			max = v
		}
	}
	return max
}

func Min(ints ...int) int {
	min := ints[0]
	for _, v := range ints {
		if min > v {
			min = v
		}
	}
	return min
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
