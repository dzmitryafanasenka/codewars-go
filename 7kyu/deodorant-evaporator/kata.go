package kata

func Evaporator(content float64, evapPerDay int, threshold int) int {
	var percentage float64 = 100
	var days int
	perDay := float64(evapPerDay)
	thold := float64(threshold)

	for percentage > thold {
		percentage = percentage - percentage*(perDay/100)
		days += 1
	}

	return days
}
