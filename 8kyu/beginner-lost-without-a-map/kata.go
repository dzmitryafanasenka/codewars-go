package kata

func Maps(x []int) []int {
	result := make([]int, 0, len(x))
	for _, v := range x {
		result = append(result, v*2)
	}

	return result
}
