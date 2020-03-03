package reverseint

import "math"

func Reverse(x int) int {
	var rev int
	for {
		pop := x % 10
		x = x / 10
		rev = (rev * 10) + pop
		if x == 0 {
			break
		}
	}
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}
	return rev
}
