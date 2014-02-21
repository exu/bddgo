package bddgo_test

// A field of N x M squares is represented by N lines of exactly M characters each. The
// character ‘*’ represents a mine and the character ‘.’ represents no-mine.

// Example input (a 3 x 4 mine-field of 12 squares, 2 of // which are mines) // 3 4
// *... ..*. ....  // Your task is to write a program to accept this input and produce as
// output a hint-field of identical dimensions where each square is a * for a mine or the
// number of adjacent mine-squares if the square does not contain a mine.

// Example output (for the above input)

// *211 12*1 0111

import (
	. "github.com/exu/bddgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Miner", func() {

	input := "3 4\n*...\n..*.\n...."
	dim := [2]int{3, 4} // todo howto converto to string?
	bombs := [][]int{
		[]int{1, 0, 0, 0},
		[]int{0, 0, 1, 0},
		[]int{0, 0, 0, 0},
	}
	hints := [][]int{
		[]int{-1, 2, 1, 1},
		[]int{1, 2, -1, 1},
		[]int{0, 1, 1, 1},
	}

	output := "*211\n12*1\n0111"

	It("detects dimensions of table", func() {
		Expect(ParseInput(input).Dimensions).To(Equal(dim))
	})

	It("it returns bombs position in array", func() {
		Expect(ParseInput(input).Bombs).To(Equal(bombs))
	})

	It("gets dimensions from string", func() {
		x, y := GetDimensions("3 4")
		Expect(x).To(Equal(3))
		Expect(y).To(Equal(4))
	})

	It("counts bombs around point", func() {
		tab := [][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{0, 1, 1},
		}
		Expect(CountBombs(tab, 1, 1)).To(Equal(7))
	})

	It("returns bombs hints", func() {
		Expect(ParseInput(input).Hints).To(Equal(hints))
	})

	It("returns output as string", func() {
		Expect(ParseInput(input).Output).To(Equal(output))
	})
})
