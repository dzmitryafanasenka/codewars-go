package kata

func Past(h, m, s int) int {
	return ((h*60+m)*60 + s) * 1000
}
