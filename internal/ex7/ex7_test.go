package ex7

import (
	"strings"
	"testing"
)

func TestCompareHands(t *testing.T) {
	assert(t, fud("22222", "33333"), -1)
	assert(t, fud("33333", "33333"), 0)
	assert(t, fud("33333", "22222"), 1)

	assert(t, fud("22222", "AAAAK"), 1)
	assert(t, fud("22223", "AAAKK"), 1)
	assert(t, fud("22233", "AAAKQ"), 1)
	assert(t, fud("22234", "AA2KK"), 1)
	assert(t, fud("22433", "AAQJT"), 1)
	assert(t, fud("22456", "AKQJT"), 1)

	assert(t, fud("65432", "23456"), 1)
	assert(t, fud("23456", "65432"), -1)
	assert(t, fud("23456", "23456"), 0)
	assert(t, fud("76543", "76542"), 1)
}

func fud(a string, b string) int {
	return compareHands(strings.Split(a, ""), strings.Split(b, ""))
}

func assert(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
