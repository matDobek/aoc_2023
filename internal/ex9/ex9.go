package ex9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1969958987
func A() {
	var sum int

	input := parseInput()

	for _, extrapolated := range extrapolate(input) {
		for _, v := range extrapolated {
			sum += v[len(v)-1]
		}
	}

	fmt.Printf("%#v", sum)
}

// 1068
func B() {
	var sum int

	input := parseInput()

	for _, extrapolated := range extrapolate(input) {
		v := extrapolated[len(extrapolated)-1][0]
		for i := len(extrapolated) - 2; i >= 0; i-- {
			a := extrapolated[i][0]
			v = a - v
		}
		sum += v
	}

	fmt.Printf("%#v", sum)
}

func extrapolate(input [][]int) [][][]int {
	result := make([][][]int, len(input))

	for ri, row := range input {
		var values []int
		values = append(values, row...)

		for !allZeros(values) {
			result[ri] = append(result[ri], values)

			var newValues []int

			for i := 0; i < len(values)-1; i++ {
				a := values[i]
				b := values[i+1]

				newValues = append(newValues, b-a)
			}

			values = newValues
		}
	}

	return result
}

func allZeros(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}

	return true
}

func parseInput() [][]int {
	input, err := os.ReadFile("./internal/ex9/ex9.input")
	if err != nil {
		panic(err)
	}

	var result [][]int
	for _, line := range strings.Split(strings.Trim(string(input), " \t\n"), "\n") {

		var set []int
		for _, str := range strings.Split(strings.Trim(line, " \t\n"), " ") {
			v, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			set = append(set, v)
		}
		result = append(result, set)
	}

	return result
}
