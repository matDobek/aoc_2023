package ex5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// 1181555926
func A() {
	input := getInput()
	seedsIn := strings.Split(input[0], " ")
	seedsIn = seedsIn[1:]
	mapsIn := input[1:]

	seeds := []int{}
	for _, s := range seedsIn {
		seed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, seed)
	}

	maps := [][][]int{}
	for _, in := range mapsIn {
		var ls [][]int
		for _, l := range strings.Split(in, "\n")[1:] {
			var nums []int
			for _, s := range strings.Split(l, " ") {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}

				nums = append(nums, num)
			}
			ls = append(ls, nums)
		}
		maps = append(maps, ls)
	}

	nums := []int{}
	for _, seed := range seeds {
		s := seed

		for _, ranges := range maps {
			for _, r := range ranges {
				min := r[1]
				max := r[1] + r[2]

				if min <= s && s < max {
					s = r[0] + (s - min)
					break
				}
			}
		}

		nums = append(nums, s)
	}

	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}

	fmt.Println(min)
}

// 37806486
func B() {
	input := getInput()
	seedsIn := strings.Split(input[0], " ")
	seedsIn = seedsIn[1:]
	mapsIn := input[1:]

	seeds := []int{}
	for _, s := range seedsIn {
		seed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, seed)
	}

	maps := [][][]int{}
	for _, in := range mapsIn {
		var ls [][]int
		for _, l := range strings.Split(in, "\n")[1:] {
			var nums []int
			for _, s := range strings.Split(l, " ") {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}

				nums = append(nums, num)
			}
			ls = append(ls, nums)
		}
		maps = append(maps, ls)
	}

	new_seeds := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seed := []int{seeds[i], seeds[i], seeds[i+1]}
		new_seeds = append(new_seeds, seed)
	}

	reversedMaps := append([][][]int{new_seeds}, maps...)
	slices.Reverse(reversedMaps)

	for num := 0; true; num++ {
		seed := num
		for i, ranges := range reversedMaps {

			for _, r := range ranges {
				min := r[0]
				max := r[0] + r[2]

				if min <= seed && seed < max {
					seed = seed - r[0] + r[1]

					if i+1 == len(reversedMaps) {
						fmt.Printf("%v: %v\n", seed, num)
						return
					} else {
						break
					}
				}
			}
		}
	}
}

func getInput() []string {
	in, err := os.ReadFile("./internal/ex5/ex5.input")
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(in), " \t\n"), "\n\n")
}
