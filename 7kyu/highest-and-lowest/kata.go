package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numbers := strings.Split(in, " ")
	low, _ := strconv.Atoi(numbers[0])
	high := low
	for _, v := range numbers {
		number, _ := strconv.Atoi(v)
		if low > number {
			low = number
		}

		if high < number {
			high = number
		}
	}

	return fmt.Sprintf("%v %v", high, low)
}
