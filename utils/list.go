package utils

func UpdateTop(list []int, val int) {
	newVal := val
	for i, v := range list {
		if newVal > v {
			list[i] = newVal
			newVal = v
		}
	}
}
