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

var _ = Describe("Should pass some basic tests", func() {
	It("Divisors(1)", func() { Expect(Divisors(1)).To(Equal(1)) })
	It("Divisors(10)", func() { Expect(Divisors(10)).To(Equal(4)) })
	It("Divisors(11)", func() { Expect(Divisors(11)).To(Equal(2)) })
	It("Divisors(54)", func() { Expect(Divisors(54)).To(Equal(8)) })
	It("Divisors(64)", func() { Expect(Divisors(64)).To(Equal(7)) })
})
