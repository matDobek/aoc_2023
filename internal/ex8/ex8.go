package ex8

import (
	"fmt"
	"os"
	"strings"
)

// 20569
func A() {
	instructions, network := parseInput()

	var inst []int
	for _, instruction := range instructions {
		if instruction == "R" {
			inst = append(inst, 1)
		} else if instruction == "L" {
			inst = append(inst, 0)
		}
	}

	currentNode := "AAA"
	keepSearching := true
	var steps int

	for keepSearching {
		for _, goTo := range inst {
			steps++
			currentNode = network[currentNode][goTo]
			if currentNode == "ZZZ" {
				keepSearching = false
				break
			}

		}
	}

	fmt.Printf("%#v\n", steps)
}

func B() {
	instructions, network := parseInput()

	var inst []int
	for _, instruction := range instructions {
		if instruction == "R" {
			inst = append(inst, 1)
		} else if instruction == "L" {
			inst = append(inst, 0)
		}
	}

	var currentNodes []string
	for node := range network {
		if strings.HasSuffix(node, "A") {
			currentNodes = append(currentNodes, node)
		}
	}

	steps := []int{}
	for _, node := range currentNodes {
		steps = append(steps, findNumberOfSteps(node, inst, network))
	}

	fmt.Printf("%#v\n", steps)
}

func findNumberOfSteps(current string, instructions []int, network map[string][]string) int {
	var steps int

	for {
		for _, goTo := range instructions {
			steps++
			current = network[current][goTo]
			if strings.HasSuffix(current, "Z") {
				return steps
			}
		}
	}
}

func parseInput() ([]string, map[string][]string) {
	input, err := os.ReadFile("./internal/ex8/ex8.input")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.Trim(string(input), " \t\n"), "\n\n")
	instructions := strings.Split(strings.Trim(parts[0], " \t\n"), "")

	network := make(map[string][]string)
	for _, line := range strings.Split(parts[1], "\n") {
		nodes := strings.Split(line, " = ")
		neighbours := strings.Split(strings.Trim(nodes[1], "()"), ", ")

		network[nodes[0]] = neighbours
	}

	return instructions, network
}
