package ex11

import (
	"fmt"
	"os"
	"strings"
)

// 9556896
func A() {
  fmt.Printf("%v\n", solve(2))
}

// 685038186836
func B() {
  fmt.Printf("%v\n", solve(1000000))
}

func solve(expansionRate int) int {
  universe := parseInput()

  var galaxies [][]int

  for i := 0; i < len(universe); i++ {
    for j := 0; j < len(universe[i]); j++ {
      if universe[i][j] == "#" {
        galaxies = append(galaxies, []int{i, j})
      }
    }
  }

  var connections [][][]int

  for _, u1 := range galaxies {
    for _, u2 := range galaxies {
      if (u2[0] > u1[0]) || (u2[0] == u1[0] && u2[1] > u1[1]) {
        pair := [][]int{
          {u1[0], u1[1]},
          {u2[0], u2[1]},
        }
        connections = append(connections, pair)
      }
    }
  }

  var emptySpacesR []int
  var emptySpacesC []int

  for i := range universe {
    allEmpty := true

    for j := range universe[0] {
      if universe[i][j] == "#" {
        allEmpty = false
        break
      }
    }

    if allEmpty {
      emptySpacesR = append(emptySpacesR, i)
    }
  }

  for i := range universe[0] {
    allEmpty := true
    for j := range universe {
      if universe[j][i] == "#" {
        allEmpty = false
        break
      }
    }

    if allEmpty {
      emptySpacesC = append(emptySpacesC, i)
    }
  }

  var sum int
  for _, c := range connections {
    a := c[0]
    b := c[1]

    diffR := b[0] - a[0]
    diffC := b[1] - a[1]

    var overheadR int
    var overheadC int

    for _, emptySpace := range emptySpacesR {

      if min(a[0], b[0]) <= emptySpace && emptySpace <= max(a[0], b[0]) {
        overheadR += expansionRate - 1
      }
    }

    for _, emptySpace := range emptySpacesC {

      if min(a[1], b[1]) <= emptySpace && emptySpace <= max(a[1], b[1]) {
        overheadC += expansionRate - 1
      }
    }

    steps := abs(diffR) + abs(diffC) + overheadR + overheadC
    sum += steps
  }

  return sum
}

func abs(x int) int {
  if x < 0 {
    return -x
  }

  return x
}

func parseInput() [][]string {
  rows, err := os.ReadFile("./internal/ex11/ex11.input")
  if err != nil {
    panic(err)
  }

  var universe [][]string
  for _, row := range strings.Split(strings.Trim(string(rows), "  \t\n"), "\n") {
    var cs []string

    for _, col := range strings.Split(row, "") {
      cs = append(cs, col)
    }

    universe = append(universe, cs)
  }

  return universe
}
