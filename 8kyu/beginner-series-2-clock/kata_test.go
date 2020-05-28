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

var _ = Describe("Basic tests", func() {
	It("Past(0, 1, 1)", func() { Expect(Past(0, 1, 1)).To(Equal(61000)) })
	It("Past(1, 1, 1)", func() { Expect(Past(1, 1, 1)).To(Equal(3661000)) })
	It("Past(0, 0, 0)", func() { Expect(Past(0, 0, 0)).To(Equal(0)) })
	It("Past(1, 0, 1)", func() { Expect(Past(1, 0, 1)).To(Equal(3601000)) })
	It("Past(1, 0, 0)", func() { Expect(Past(1, 0, 0)).To(Equal(3600000)) })
})
