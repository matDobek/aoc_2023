package ex6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 741000
func A() {
	times, distances := parseInput()
	res := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		dist := distances[i]

		succ := 0
		for t := 0; t <= time; t++ {
			d := (time - t) * t
			if d > dist {
				succ++
			}
		}
		res *= succ
	}

	fmt.Println(res)
}

// 38220708
func B() {
	timesIn, distancesIn := parseInput()

	timeStr := ""
	for _, str := range timesIn {
		timeStr += fmt.Sprintf("%d", str)
	}

	distanceStr := ""
	for _, str := range distancesIn {
		distanceStr += fmt.Sprintf("%d", str)
	}

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(err)
	}

	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		panic(err)
	}

	succ := 0
	for t := 0; t <= time; t++ {
		d := (time - t) * t
		if d > distance {
			succ++
		}
	}
	fmt.Println(succ)
}

func parseInput() ([]int, []int) {
	f, err := os.ReadFile("./internal/ex6/ex6.input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.Trim(string(f), " \t\n"), "\n")
	var times []int
	var distances []int

	for _, v := range strings.Split(input[0], " ")[1:] {
		str := strings.Trim(v, " \t\n")
		if str == "" {
			continue
		}

		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		times = append(times, n)
	}

	for _, v := range strings.Split(input[1], " ")[1:] {
		str := strings.Trim(v, " \t\n")
		if str == "" {
			continue
		}

		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		distances = append(distances, n)
	}

	return times, distances
}
