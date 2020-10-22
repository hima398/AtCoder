package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadNumber0(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		".#.#.",
		".#.#.",
		".#.#.",
		".###.",
	}

	assert.Equal(t, "0", ReadNumber(lines, w))
}

func TestReadNumber1(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		"..#..",
		".##..",
		"..#..",
		"..#..",
		".###.",
	}

	assert.Equal(t, "1", ReadNumber(lines, w))
}

func TestReadNumber2(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		"...#.",
		".###.",
		".#...",
		".###.",
	}

	assert.Equal(t, "2", ReadNumber(lines, w))
}
func TestReadNumber3(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		"...#.",
		".###.",
		"...#.",
		".###.",
	}

	assert.Equal(t, "3", ReadNumber(lines, w))
}
func TestReadNumber4(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".#.#.",
		".#.#.",
		".###.",
		"...#.",
		"...#.",
	}

	assert.Equal(t, "4", ReadNumber(lines, w))
}
func TestReadNumber5(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		".#...",
		".###.",
		"...#.",
		".###.",
	}

	assert.Equal(t, "5", ReadNumber(lines, w))
}
func TestReadNumber6(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		".#...",
		".###.",
		".#.#.",
		".###.",
	}

	assert.Equal(t, "6", ReadNumber(lines, w))
}
func TestReadNumber7(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		"...#.",
		"...#.",
		"...#.",
		"...#.",
	}

	assert.Equal(t, "7", ReadNumber(lines, w))
}
func TestReadNumber8(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		".#.#.",
		".###.",
		".#.#.",
		".###.",
	}

	assert.Equal(t, "8", ReadNumber(lines, w))
}
func TestReadNumber9(t *testing.T) {
	n := 1
	w := 4*n + 1
	lines := []string{
		".###.",
		".#.#.",
		".###.",
		"...#.",
		".###.",
	}

	assert.Equal(t, "9", ReadNumber(lines, w))
}

func TestReadNumberInput1(t *testing.T) {
	n := 10
	w := 4*n + 1
	lines := []string{
		".###..#..###.###.#.#.###.###.###.###.###.",
		".#.#.##....#...#.#.#.#...#.....#.#.#.#.#.",
		".#.#..#..###.###.###.###.###...#.###.###.",
		".#.#..#..#.....#...#...#.#.#...#.#.#...#.",
		".###.###.###.###...#.###.###...#.###.###.",
	}
	assert.Equal(t, "0123456789", ReadNumber(lines, w))
}

func TestReadNumberInput2(t *testing.T) {
	n := 29
	w := 4*n + 1
	lines := []string{
		".###.###.###.###.###.###.###.###.###.#.#.###.#.#.#.#.#.#.###.###.###.###..#..###.###.###.###.###.#.#.###.###.###.###.",
		"...#.#.#...#.#.#.#.#.#...#.#...#.#.#.#.#.#...#.#.#.#.#.#.#.....#.#.#.#.#.##..#.#...#.#.#...#.#...#.#...#.#.....#...#.",
		".###.#.#...#.###.#.#.###.###...#.###.###.###.###.###.###.###...#.###.#.#..#..###...#.###.###.###.###.###.###.###.###.",
		".#...#.#...#...#.#.#.#.#...#...#.#.#...#.#.#...#...#...#.#.#...#...#.#.#..#..#.#...#...#.#...#.#...#.#.....#...#.#...",
		".###.###...#.###.###.###.###...#.###...#.###...#...#...#.###...#.###.###.###.###...#.###.###.###...#.###.###.###.###."}
	assert.Equal(t, "20790697846444679018792642532", ReadNumber(lines, w))
}
