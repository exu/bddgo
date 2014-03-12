package bddgo

// import "fmt"

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

var divisors = map[int]string{
	1000:          "thousand",
	1000000:       "million",
	1000000000:    "billion",
	1000000000000: "billiard",
}

func isPredefined(number int) bool {
	_, ok := predefined[number]
	return ok
}

func getS(number int) string {
	if number <= 1 {
		return ""
	}

	return "s"
}

func findDivisor(number int) int {
	divisor := 1
	for {
		divisor *= 1000
		if divisor > number {
			return divisor / 1000
		}
	}
}

func divideNumber(number int) map[int]int {
	result := map[int]int{}
	divisor := findDivisor(number)

	for {
		if divisor == 0 {
			return result
		}
		result[divisor] = number / divisor
		number = number - (divisor * result[divisor])
		divisor = divisor / 1000
	}

	return result
}

func spellBasic(number int) string {
	return predefined[number]
}

func spellHundrets(number int) string {
	hundreds := number / 100
	if number == 0 {
		return ""
	}

	spelled := ""
	if hundreds > 0 {
		spelled = predefined[hundreds] + " hundred and "
	}

	return spelled + spellTens(number%100)
}

func spellTens(number int) string {

	if isPredefined(number) {
		return predefined[number]
	}

	spelled := ""
	ten := 10 * (number / 10)
	one := number % 10

	if ten > 0 {
		spelled += predefined[ten]
	}

	if ten > 0 && one > 0 {
		spelled += " "
	}

	if one > 0 {
		spelled += predefined[one]
	}

	return spelled
}

func Spell(number int) string {
	if isPredefined(number) {
		return predefined[number]
	}

	if number < 100 {
		return spellTens(number)
	}

	spelled := ""

	divided := divideNumber(number)
	s, div, sub := "", "", ""

	for multiplier, num := range divided {
		if multiplier > 1 {
			div = " " + divisors[multiplier]
			s = getS(num) + " "
		} else {
			s, div = "", ""
		}

		if isPredefined(num) {
			sub = spellBasic(num)
		} else {
			sub = spellHundrets(num)
		}

		spelled += sub + div + s
	}

	return spelled
}
