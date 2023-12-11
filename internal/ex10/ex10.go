package ex10

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 6754
func A() {
	start, rows := parseInput()

	starters := findConnectedNeighbors(rows, start)

	currA := start
	currB := start
	nextA := starters[0]
	nextB := starters[1]

	for i := 0; true; i++ {
		if reflect.DeepEqual(nextA, nextB) {
			fmt.Printf("%v\n", i+1)
			break
		}

		tmpA := nextA
		tmpB := nextB

		nextA = nextNeighbour(rows, currA, nextA)
		nextB = nextNeighbour(rows, currB, nextB)

		currA = tmpA
		currB = tmpB
	}
}

// 567
func B() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	start, board := parseInput()

	edges := [][]string{
		{"s", "s"},
		{"u", "d"},
		{"d", "l"},
		{"d", "r"},
	}

	isMainPipe := make([][]bool, len(board))
	isInside := make([][]bool, len(board))

	for i := 0; i < len(board); i++ {
		isMainPipe[i] = make([]bool, len(board[0]))
		isInside[i] = make([]bool, len(board[0]))
	}

	starters := findConnectedNeighbors(board, start)

	curr := start
	next := starters[0]

	for {
		isMainPipe[curr[0]][curr[1]] = true
		showBoard(board, isMainPipe, isInside)

		if len(next) == 0 {
			break
		}

		tmp := next
		next = nextNeighbour(board, curr, next)
		curr = tmp
	}

	for i := range isMainPipe {
		var in bool

		for j := range isMainPipe[i] {

			for _, edge := range edges {
				if isMainPipe[i][j] && reflect.DeepEqual(board[i][j], edge) {
					in = !in
				}
			}

			if !isMainPipe[i][j] && in { // if not a pipe and inside
				isInside[i][j] = true
				showBoard(board, isMainPipe, isInside)
			}
		}
	}

	var counter int
	for i := range isInside {
		for j := range isInside[i] {
			if isInside[i][j] {
				counter++
			}
		}
	}

	// fmt.Println(counter)
}

func showBoard(board [][][]string, pipes [][]bool, inside [][]bool) {
	colorPipe := "\033[0;32m"
	colorInside := "\033[0;31m"
	colorNone := "\033[0m"

	time.Sleep(100 * time.Millisecond)

	l := strconv.Itoa(len(board) * 2)
	fmt.Printf("\033[" + l + "A")

	for i, row := range board {
		for j, col := range row {
			if inside[i][j] && pipes[i][j] {
				fmt.Printf(colorInside)
			} else if pipes[i][j] {
				fmt.Printf(colorPipe)
			} else if inside[i][j] {
				fmt.Printf(colorInside)
			}

			if reflect.DeepEqual(col, []string{".", "."}) {
				fmt.Printf(".")
			} else if reflect.DeepEqual(col, []string{"s", "s"}) {
				fmt.Printf("s")
			} else if reflect.DeepEqual(col, []string{"u", "d"}) {
				fmt.Printf("│")
			} else if reflect.DeepEqual(col, []string{"l", "r"}) {
				fmt.Printf("─")
			} else if reflect.DeepEqual(col, []string{"d", "l"}) {
				fmt.Printf("┐")
			} else if reflect.DeepEqual(col, []string{"d", "r"}) {
				fmt.Printf("┌")
			} else if reflect.DeepEqual(col, []string{"u", "l"}) {
				fmt.Printf("┘")
			} else if reflect.DeepEqual(col, []string{"u", "r"}) {
				fmt.Printf("└")
			}

			fmt.Printf(colorNone)
		}
		fmt.Println()
	}

}

func findConnectedNeighbors(board [][][]string, curr []int) [][]int {
	result := [][]int{}
	maybeValid := [][]int{
		{curr[0] - 1, curr[1]},
		{curr[0] + 1, curr[1]},
		{curr[0], curr[1] - 1},
		{curr[0], curr[1] + 1},
	}

	for _, maybe := range maybeValid {
		if withinBounds(board, maybe) &&
			connected(board, curr, maybe) {
			result = append(result, maybe)
		}
	}

	return result
}

func withinBounds(board [][][]string, a []int) bool {
	maxR := len(board) - 1
	maxC := len(board[0]) - 1

	if a[0] < 0 || a[0] > maxR {
		return false
	}

	if a[1] < 0 || a[1] > maxC {
		return false
	}

	return true
}

func connected(board [][][]string, a []int, b []int) bool {
	diffR := b[0] - a[0]
	diffC := b[1] - a[1]

	var aMustInclude string
	var bMustInclude string
	if reflect.DeepEqual([]int{diffR, diffC}, []int{-1, 0}) {
		aMustInclude = "u"
		bMustInclude = "d"
	}
	if reflect.DeepEqual([]int{diffR, diffC}, []int{1, 0}) {
		aMustInclude = "d"
		bMustInclude = "u"
	}
	if reflect.DeepEqual([]int{diffR, diffC}, []int{0, -1}) {
		aMustInclude = "l"
		bMustInclude = "r"
	}
	if reflect.DeepEqual([]int{diffR, diffC}, []int{0, 1}) {
		aMustInclude = "r"
		bMustInclude = "l"
	}

	var aIncludes bool
	var bIncludes bool

	for _, a := range board[a[0]][a[1]] {
		if a == aMustInclude || a == "s" {
			aIncludes = true
		}
	}

	for _, b := range board[b[0]][b[1]] {
		if b == "s" {
			bIncludes = false
		}

		if b == bMustInclude {
			bIncludes = true
		}
	}

	return aIncludes && bIncludes
}

func nextNeighbour(board [][][]string, prev, curr []int) []int {
	var next []int

	connected := findConnectedNeighbors(board, curr)
	for _, c := range connected {
		if !reflect.DeepEqual(c, prev) {
			next = c
		}
	}

	return next
}

func parseInput() ([]int, [][][]string) {
	var result [][][]string
	var start []int

	input, err := os.ReadFile("./internal/ex10/ex10.input")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.Trim(string(input), " \t\n"), "\n")

	for i, row := range rows {
		result = append(result, [][]string{})

		for j, s := range strings.Split(row, "") {
			var directions []string
			switch s {
			case "S":
				start = []int{i, j}
				directions = []string{"s", "s"}
			case ".":
				directions = []string{".", "."}
			case "|":
				directions = []string{"u", "d"}
			case "-":
				directions = []string{"l", "r"}
			case "7":
				directions = []string{"d", "l"}
			case "F":
				directions = []string{"d", "r"}
			case "J":
				directions = []string{"u", "l"}
			case "L":
				directions = []string{"u", "r"}
			}
			result[i] = append(result[i], directions)
		}
	}

	return start, result
}
