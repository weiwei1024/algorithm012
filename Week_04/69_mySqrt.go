package src

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)/2
		n := mid * mid
		if n == x {
			return n
		} else if n < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}
