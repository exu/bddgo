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

type TestProvider struct {
	input      string
	dimensions [2]string
	bombs      map[int][]int
	hints      map[int][]int
	output     string
}

var data TestProvider

var _ = Describe("Miner", func() {

	data.input = "3 4\n*...\n..*.\n...."
	data.dimensions = [2]string{"3", "4"} // todo howto converto to string?
	data.bombs = map[int][]int{
		0: []int{1, 0, 0, 0},
		1: []int{0, 0, 1, 0},
		2: []int{0, 0, 0, 0},
	}
	data.hints = map[int][]int{
		0: []int{-1, 2, 1, 1},
		1: []int{1, 2, -1, 1},
		2: []int{0, 1, 1, 1},
	}
	data.output = "*211\n12*1\n0111"

	It("detects dimensions of table", func() {
		Expect(ParseInput(data.input).dimensions).To(Equal(data.dimensions))
	})

	It("gives us bombs positions from string", func() {
		// Expect(FindBombs(data.input)).To(Equal(data.bombs))
	})

	It("gives us bombs positions from string", func() {
		// Expect(GetHints(data.input)).To(Equal(data.hints))
	})

	It("gives us bombs positions from string", func() {
		// Expect(Play(data.input)).To(Equal(data.hints))
	})
})
