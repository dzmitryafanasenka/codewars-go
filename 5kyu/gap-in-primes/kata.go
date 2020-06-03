package kata

import (
	"math"
)

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}

	return true
}

func Gap(g, m, n int) []int {
	if n < 2 || m < 2 || g > n-m || n < m {
		return nil
	}

	latest := math.MinInt32

	for current := m; current < n; current++ {
		if IsPrime(current) {
			if current-latest == g {
				return []int{latest, current}
			} else {
				latest = current
			}
		}
	}

	return nil
}
