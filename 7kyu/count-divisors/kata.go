package kata

import "math"

func Divisors(n int) int {
	var count int

	for i := 1; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			if n/i == i {
				count++
			} else {
				count += 2
			}
		}
	}

	return count
}
