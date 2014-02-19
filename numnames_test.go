package bddgo_test

import (
	. "github.com/exu/bddgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var teens = []int{10, 11, 12, 13, 14, 15, 16, 18, 19}

var tens = map[int]string{
	22: "twenty two",
	31: "thirty one",
	43: "fourty three",
	56: "fifty six",
	64: "sixty four",
	78: "seventy eight",
	85: "eighty five",
	98: "ninety eight",
	99: "ninety nine",
}

var hundreds = map[int]string{
	544: "five hundred and fourty four",
	987: "nine hundred and eighty seven",
	110: "one hundred and ten",
	1:   "one",
	13:  "thirteen",
}

var tousands = map[int]string{
	898982332876: "eight hundred and ninety eight billions nine hundred and eighty two millions three hundred and thirty two thousands eight hundred and seventy six",
	32876:        "thirty two thousands eight hundred and seventy six",
	1544:         "one thousand five hundred and fourty four",
	2987:         "two thousands nine hundred and eighty seven",
	3110:         "three thousands one hundred and ten",
}

var predefined = map[int]string{
	0:   "zero",
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	10:  "ten",
	11:  "eleven",
	12:  "twelve",
	13:  "thirteen",
	14:  "fourteen",
	15:  "fifteen",
	16:  "sixteen",
	17:  "seventeen",
	18:  "eighteen",
	19:  "nineteen",
	20:  "twenty",
	30:  "thirty",
	40:  "fourty",
	50:  "fifty",
	60:  "sixty",
	70:  "seventy",
	80:  "eighty",
	90:  "ninety",
	100: "hundred",
}

var _ = Describe("Numnames", func() {

	It("Spells ones", func() {
		for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 8, 9} {
			Expect(Spell(num)).To(Equal(predefined[num]))
		}
	})

	It("Spells ones", func() {
		for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 8, 9} {
			Expect(Spell(num)).To(Equal(predefined[num]))
		}
	})

	It("Spells teens", func() {
		for _, num := range teens {
			Expect(Spell(num)).To(Equal(predefined[num]))
		}
	})

	It("Spells tens", func() {
		for k, v := range tens {
			Expect(Spell(k)).To(Equal(v))
		}
	})

	It("Spells hundreds", func() {
		for num, text := range hundreds {
			Expect(Spell(num)).To(Equal(text))
		}
	})

	It("Spells tousands", func() {
		for num, text := range tousands {
			Expect(Spell(num)).To(Equal(text))
		}
	})
})
