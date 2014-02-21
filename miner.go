package bddgo

import (
	"bytes"
	"strconv"
	"strings"
)

type Output struct {
	Dimensions [2]int
	Bombs      [][]int
	Hints      [][]int
	Output     string
}

func ParseInput(input string) Output {
	var output Output
	var os bytes.Buffer

	data := strings.Split(input, "\n")
	for k, v := range data {
		if k == 0 {
			x, y := GetDimensions(data[0])

			output.Dimensions = [2]int{x, y}

			output.Hints = make([][]int, x)
			for k, _ := range output.Hints {
				output.Hints[k] = make([]int, y)
			}
		} else {
			cells := strings.Split(v, "")
			row := []int{}
			for _, v := range cells {
				if v == "*" {
					row = append(row, 1)
				} else {
					row = append(row, 0)
				}
			}

			output.Bombs = append(output.Bombs, row)
		}
	}

	for i, a := range output.Bombs {
		for j, v := range a {
			if v == 0 {
				output.Hints[i][j] = CountBombs(output.Bombs, i, j)
				os.WriteString(strconv.Itoa(output.Hints[i][j]))
			} else {
				output.Hints[i][j] = -1
				os.WriteString("*")
			}
		}

		if i < len(output.Bombs)-1 {
			os.WriteString("\n")
		}
	}
	output.Output = os.String()

	return output
}

func GetDimensions(str string) (x int, y int) {
	dims := strings.Split(str, " ")
	x, _ = strconv.Atoi(dims[0])
	y, _ = strconv.Atoi(dims[1])

	return x, y
}

func CountBombs(arr [][]int, i int, j int) int {
	directions := [][]int{
		[]int{0, -1},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, -1},
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
	}

	count := 0
	for _, d := range directions {
		if i+d[0] >= 0 && i+d[0] < len(arr) {
			if j+d[1] >= 0 && j+d[1] < len(arr[i+d[0]]) {
				count += arr[i+d[0]][j+d[1]]
			}
		}
	}

	return count
}
