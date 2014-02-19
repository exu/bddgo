package bddgo_test

import (
	. "github.com/exu/bddgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Your task is to process a sequence of integer numbers
// to determine the following statistics:

// o) minimum value
// o) maximum value
// o) number of elements in the sequence
// o) average value

// For example: [6, 9, 15, -2, 92, 11]

// o) minimum value = -2
// o) maximum value = 92
// o) number of elements in the sequence = 6
// o) average value = 21.833333

var _ = Describe("Stats generator", func() {

	var result map[string]int

	BeforeEach(func() {
		input := []int{3, 3, 4, 5, 6, 99, -99, 7, 8, 64}
		result = Collect(input)
	})

	It("collects min information", func() {
		Expect(result["min"]).To(Equal(-99))
	})
	It("collects max information", func() {
		Expect(result["max"]).To(Equal(99))
	})
	It("collects avg information", func() {
		Expect(result["avg"]).To(Equal(10))
	})
	It("collects len information", func() {
		Expect(result["len"]).To(Equal(10))
	})
})
