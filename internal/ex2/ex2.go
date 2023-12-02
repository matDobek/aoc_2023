package ex2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 2377
func A() {
	input := getInput()

	sum := 0
	box1 := map[string]int{"red": 12, "green": 13, "blue": 14}
	for i, ix := range input {
		line := strings.Split(ix, ":")

		b := true
		for _, jx := range strings.Split(line[1], ";") {
			box2 := map[string]int{"red": 0, "green": 0, "blue": 0}
			for _, zx := range strings.Split(jx, ",") {
				zx_f := strings.Trim(zx, " ")
				crystal := strings.Split(zx_f, " ")

				num, err := strconv.Atoi(crystal[0])
				if err != nil { panic(err) }
				box2[crystal[1]] += num
			}

			for k, _ := range box1 {
				if box2[k] > box1[k] {
					b = false
				}
			}
		}

		if b { sum += i+1 }
	}

	fmt.Println(sum)
}

// 71220
func B() {
	input := getInput()

	sum := 0
	for _, ix := range input {
		line := strings.Split(ix, ":")

		b := true
		box1 := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, jx := range strings.Split(line[1], ";") {
			box2 := map[string]int{"red": 0, "green": 0, "blue": 0}

			for _, zx := range strings.Split(jx, ",") {
				zx_f := strings.Trim(zx, " ")
				crystal := strings.Split(zx_f, " ")

				num, err := strconv.Atoi(crystal[0])
				if err != nil { panic(err) }
				box2[crystal[1]] += num
			}

			for k, _ := range box1 {
				if box2[k] > box1[k] {
					box1[k] = box2[k]
				}
			}
		}

		pow := box1["red"] * box1["green"] * box1["blue"]
		if b { sum += pow }
	}

	fmt.Println(sum)
}

func getInput() []string {
	in, err := os.ReadFile("./internal/ex2/ex2.input")
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(in), " \t\n"), "\n")
}
