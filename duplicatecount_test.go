package kata_test

import (
	kata "kata/duplicate_count"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(prod string, exp int) {
	ans := kata.Duplicate_count(prod)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest("abcde", 0)
		dotest("abcdea", 1)
		dotest("abcdeaB11", 3)
		dotest("indivisibility", 1)
	})
})
