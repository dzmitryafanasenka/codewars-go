package kata

import (
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	var latestConsec string
	var longest string

	for i := 0; i <= len(strarr)-k; i++ {
		latestConsec = strings.Join(strarr[i:i+k], "")
		if len(latestConsec) > len(longest) {
			longest = latestConsec
		}
	}

	return longest
}
