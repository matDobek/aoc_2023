package ex7

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

const (
	Nothing = iota
	OnePair
	TwoPairs
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var cardStrength = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func handType(hand []string) int {
	m := make(map[string]int)
	for i := 0; i < len(hand); i++ {
		m[hand[i]]++
	}

	values := []int{}
	for k, v := range m {
		if k != "J" {
			values = append(values, v)
		}
	}
	slices.Sort(values)

	if len(values) == 0 {
		values = append(values, 5)
	} else {
		values[len(values)-1] += m["J"]
	}

	if reflect.DeepEqual(values, []int{5}) {
		return FiveOfKind
	}
	if reflect.DeepEqual(values, []int{1, 4}) {
		return FourOfKind
	}
	if reflect.DeepEqual(values, []int{2, 3}) {
		return FullHouse
	}
	if reflect.DeepEqual(values, []int{1, 1, 3}) {
		return ThreeOfKind
	}
	if reflect.DeepEqual(values, []int{1, 2, 2}) {
		return TwoPairs
	}
	if reflect.DeepEqual(values, []int{1, 1, 1, 2}) {
		return OnePair
	}
	if reflect.DeepEqual(values, []int{1, 1, 1, 1, 1}) {
		return Nothing
	}

	panic("unexpected hand")
}

func compareHands(aHand []string, bHand []string) int {
	aType := handType(aHand)
	bType := handType(bHand)

	if aType > bType {
		return 1
	}

	if aType < bType {
		return -1
	}

	// check order if the same hand
	for i := 0; i < len(aHand); i++ {
		if cardStrength[aHand[i]] > cardStrength[bHand[i]] {
			return 1
		}

		if cardStrength[aHand[i]] < cardStrength[bHand[i]] {
			return -1
		}
	}

	return 0
}

// 250347426
func A() {
	// m := parseInput()
	//
	// var hands [][]string
	// for hand := range m {
	// 	xs := strings.Split(hand, "")
	// 	hands = append(hands, xs)
	// }
	//
	// slices.SortStableFunc(hands, compareHands)
	//
	// var sum int
	// for i, hand := range hands {
	// 	bid := m[strings.Join(hand, "")]
	// 	sum += (i + 1) * bid
	// }
	//
	// fmt.Println(sum)
}

func B() {
	m := parseInput()

	var hands [][]string
	for hand := range m {
		xs := strings.Split(hand, "")
		hands = append(hands, xs)
	}

	slices.SortStableFunc(hands, compareHands)

	var sum int
	for i, hand := range hands {
		bid := m[strings.Join(hand, "")]
		// aType := handType(hand)
		// fmt.Printf("%v -- %v -- %v: %v\n", hand, aType, i+1, bid)
		sum += (i + 1) * bid
	}

	fmt.Println(sum)
}

func parseInput() map[string]int {
	input, err := os.ReadFile("./internal/ex7/ex7.input")
	if err != nil {
		panic(err)
	}

	m := make(map[string]int)
	for _, l := range strings.Split(strings.Trim(string(input), " \n"), "\n") {
		vals := strings.Split(strings.Trim(l, " "), " ")

		bet, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		m[vals[0]] = bet
	}

	return m
}
