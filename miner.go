package bddgo

import (
	"fmt"
	"strings"
)

type Output struct {
	dimensions []string
	bombs      map[int][]int
}

func ParseInput(input string) Output {
	var output Output

	data := strings.Split(input, "\n")

	for k, v := range data {
		if k == 0 {
			output.dimensions = strings.Split(data[0], " ")
		} else {
			cells := strings.Split(v, "")
			fmt.Println(cells)
		}
	}

	return output
}
