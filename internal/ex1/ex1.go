package ex1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 55029
func A() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	rx, _ := regexp.Compile("\\d")
	for _, i := range input {
		o := rx.FindAllString(i, -1)

		first := o[0]
		last := o[len(o)-1]
		val, _ := strconv.Atoi(first + last)
		sum += val
	}

	fmt.Printf("%#v", sum)
}

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// 55686
func B() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, in := range input {
		f := to_i(first(in))
		l := to_i(last(in))

		sum += f*10 + l
	}

	fmt.Printf("%#v", sum)
}

func first(in string) string {
	str := ""

	for i := 0; i < len(in); i++ {
		if '0' <= in[i] && in[i] <= '9' {
			return string(in[i])
		}

		for j := 0; j < len(numbers); j++ {
			if strings.HasSuffix(in[:i+1], numbers[j]) {
				return numbers[j]
			}
		}
	}

	return str
}

func last(in string) string {
	for i := len(in) - 1; i >= 0; i-- {
		if '0' <= in[i] && in[i] <= '9' {
			return string(in[i])
		}

		for j := 0; j < len(numbers); j++ {
			if strings.HasSuffix(in[:i+1], numbers[j]) {
				return numbers[j]
			}
		}
	}

	return ""
}

func to_i(str string) int {
	for i, s := range numbers {
		if str == s {
			return i + 1
		}
	}

	x, _ := strconv.Atoi(str)
	return x
}

func parseInput() ([]string, error) {
	dat, err := os.ReadFile("./internal/ex1/ex1.input")
	if err != nil {
		return nil, err
	}

	r1 := strings.Trim(string(dat), " \t\n")
	r2 := strings.Split(r1, "\n")

	return r2, nil
}
