package kata_test

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("should work for sample test cases", func() {
		Expect(kata.ReverseWords("The quick brown fox jumps over the lazy dog.")).To(Equal("ehT kciuq nworb xof spmuj revo eht yzal .god"))
		Expect(kata.ReverseWords("apple")).To(Equal("elppa"))
		Expect(kata.ReverseWords("a b c d")).To(Equal("a b c d"))
		Expect(kata.ReverseWords("double  spaced  words")).To(Equal("elbuod  decaps  sdrow"))
		Expect(kata.ReverseWords("stressed desserts")).To(Equal("desserts stressed"))
		Expect(kata.ReverseWords("hello hello")).To(Equal("olleh olleh"))
	})
})
