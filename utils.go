package leetcomm

func Max(args ...int) int {
	max := args[0]
	for i := 1; i < len(args); i++ {
		if max < args[i] {
			max = args[i]
		}
	}
	return max
}

func Min(args ...int) int {
	min := args[0]
	for i := 1; i < len(args); i++ {
		if min > args[i] {
			min = args[i]
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
