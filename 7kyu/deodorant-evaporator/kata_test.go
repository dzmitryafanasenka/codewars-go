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

func dotest(content float64, evapPerDay int, threshold int, exp int) {
	var ans = Evaporator(content, evapPerDay, threshold)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Evaporator", func() {

	It("should handle basic cases", func() {
		dotest(10, 10, 10, 22)
		dotest(10, 10, 5, 29)
		dotest(100, 5, 5, 59)

	})
})
