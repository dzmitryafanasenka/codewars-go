package kata

func Add(n int) func(int) int {
	return func(i int) int {
		return n + i
	}
}
