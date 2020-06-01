package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

func dotest(strarr []string, k int, exp string) {
	var ans = LongestConsec(strarr, k)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 0, "")
		dotest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta")
		dotest([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
			"oocccffuucccjjjkkkjyyyeehh")
		dotest([]string{}, 3, "")
		dotest([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	})
})
