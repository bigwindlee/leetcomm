package leetcomm

func Max(a int, args ...int) int {
	max := a
	for _, val := range args {
		if max < val {
			max = val
		}
	}
	return max
}

func Min(a int, args ...int) int {
	min := a
	for _, val := range args {
		if min > val {
			min = val
		}
	}
	return min
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
