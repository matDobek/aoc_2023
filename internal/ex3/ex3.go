package ex3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 560670
func A() {
	input := [][]byte{}
	for _, line := range getInput() {
		input = append(input, []byte(line))
	}

	var sum int
	var num []byte
	var isPartNum bool

	for r, row := range input {
		for c, col := range row {
			if '0' <= col && col <= '9' {
				num = append(num, col)

				for _, rr := range []int{r - 1, r, r + 1} {
					for _, cc := range []int{c - 1, c, c + 1} {
						if 0 <= rr && rr < len(input) && 0 <= cc && cc < len(input[rr]) {
							if (input[rr][cc] < '0' || '9' < input[rr][cc]) && input[rr][cc] != '.' {
								isPartNum = true
							}
						}
					}
				}
			} else {
				if isPartNum {
					v, err := strconv.Atoi(string(num))
					if err != nil {
						panic(err)
					}

					sum += v
				}

				num = []byte{}
				isPartNum = false
			}
		}
	}

	fmt.Println(sum)
}

// 91622824
func B() {
	input := [][]byte{}
	for _, line := range getInput() {
		input = append(input, []byte(line))
	}

	var num []byte
	var isPartNum bool
	var locs [][2]int
	m := make(map[[2]int][]int)

	for r, row := range input {
		for c, col := range row {
			if '0' <= col && col <= '9' {
				num = append(num, col)

				for _, rr := range []int{r - 1, r, r + 1} {
					for _, cc := range []int{c - 1, c, c + 1} {
						if 0 <= rr && rr < len(input) && 0 <= cc && cc < len(input[rr]) {
							if input[rr][cc] == '*' {
								isPartNum = true
								locs = append(locs, [2]int{rr, cc})
							}
						}
					}
				}

			} else {
				if isPartNum {
					v, err := strconv.Atoi(string(num))
					if err != nil {
						panic(err)
					}

					uniq := make(map[[2]int]struct{})
					for _, l := range locs {
						uniq[l] = struct{}{}
					}
					for k := range uniq {
						m[k] = append(m[k], v)
					}
				}

				num = []byte{}
				locs = [][2]int{}
				isPartNum = false
			}
		}
	}

	sum := 0
	for _, v := range m {
		if len(v) < 2 {
			continue
		}

		mult := 1
		for _, vv := range v {
			mult *= vv
		}

		sum += mult
	}
	fmt.Println(sum)
}

func getInput() []string {
	input, err := os.ReadFile("./internal/ex3/ex3.input")
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(input), " \t\n"), "\n")
}
