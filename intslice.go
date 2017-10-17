package leetcomm

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func IntSliceEqualIgnoreOrder(a, b []int) bool {
	xor, suma, sumb := 0, 0, 0
	for _, val := range a {
		xor ^= val
		suma += val * val
	}
	for _, val := range b {
		xor ^= val
		sumb += val * val
	}
	return xor == 0 && suma == sumb
}

func Int2DSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func DupSlice(a []int) []int {
	ret := make([]int, len(a))
	copy(ret, a)
	return ret
}

func ReverseSlice(a []int) {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-1-i] = a[l-1-i], a[i]
	}
}
