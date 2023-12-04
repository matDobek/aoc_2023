package ex4

import (
	"fmt"
	"os"
	"strings"
)

// 23441
func A() {
	lines := getInput()
	sum := 0

	for _, line := range lines {
		ll := strings.Split(line, ":")
		ll = strings.Split(ll[1], "|")

		winning := []string{}
		for _, v := range strings.Split(strings.Trim(ll[0], " "), " ") {
			v = strings.Trim(v, " ")
			if v != "" {
				winning = append(winning, v)
			}
		}

		nums := []string{}
		for _, v := range strings.Split(strings.Trim(ll[1], " "), " ") {
			v = strings.Trim(v, " ")
			if v != "" {
				nums = append(nums, v)
			}
		}

		claimed := []bool{}

		for range winning {
			claimed = append(claimed, false)
		}

		var streak int
		for _, num := range nums {
			for j, winNum := range winning {
				if num != winNum {
					continue
				}

				if claimed[j] {
					continue
				}

				if streak == 0 {
					streak = 1
				} else {
					streak = streak * 2
				}
				claimed[j] = true
				break
			}
		}

		sum += streak
	}

	fmt.Println(sum)
}

// 5923918
func B() {
	lines := getInput()
	sum := 0

	cards := make(map[int]int)
	for i := range lines {
		cards[i] = 1
	}

	for i, line := range lines {
		ll := strings.Split(line, ":")
		ll = strings.Split(ll[1], "|")

		winning := []string{}
		for _, v := range strings.Split(strings.Trim(ll[0], " "), " ") {
			v = strings.Trim(v, " ")
			if v != "" {
				winning = append(winning, v)
			}
		}

		nums := []string{}
		for _, v := range strings.Split(strings.Trim(ll[1], " "), " ") {
			v = strings.Trim(v, " ")
			if v != "" {
				nums = append(nums, v)
			}
		}

		claimed := []bool{}
		for range winning {
			claimed = append(claimed, false)
		}

		var streak int
		for _, num := range nums {
			for z, winNum := range winning {
				if num != winNum {
					continue
				}

				if claimed[z] {
					continue
				}

				streak++
				claimed[z] = true
				break
			}
		}

		for j := i + 1; j < i+streak+1; j++ {
			cards[j] += cards[i]
		}
	}

	for _, v := range cards {
		sum += v
	}

	fmt.Println(sum)
}

func getInput() []string {
	input, err := os.ReadFile("./internal/ex4/ex4.input")
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(input), " \t\n"), "\n")
}
