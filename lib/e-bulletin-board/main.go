package main

import "fmt"

//.### ..#. .### .### .#.# .### .### .### .### .###.
//.#.# .##. ...# ...# .#.# .#.. .#.. ...# .#.# .#.#.
//.#.# ..#. .### .### .### .### .### ...# .### .###.
//.#.# ..#. .#.. ...# ...# ...# .#.# ...# .#.# ...#.
//.### .### .### .### ...# .### .### ...# .### .###.

//.###..#..###.###.#.#.###.###.###.###.###.
//.#.#.##....#...#.#.#.#...#.....#.#.#.#.#.
//.#.#..#..###.###.###.###.###...#.###.###.
//.#.#..#..#.....#...#...#.#.#...#.#.#...#.
//.###.###.###.###...#.###.###...#.###.###.

var cMap = map[rune]int{
	'.': 0,
	'#': 1,
}
var m = map[int]int{
	0x7555: 0,
	0x2622: 1,
	0x7174: 2,
	0x7171: 3,
	0x5571: 4,
	0x7471: 5,
	0x7475: 6,
	0x7111: 7,
	0x7575: 8,
	0x7571: 9,
}

func ParseNumber(s string) int {
	var v = 0
	for _, c := range s {
		v |= cMap[c]
		v = v << 1
	}
	v = v >> 1
	fmt.Printf("%x\n", v)
	return m[v]
}

func ReadNumber(lines []string, n int) int {
	var v = 0
	for i := 0; i < n-1; i += 4 {
		s := lines[0][i : i+4]
		s += lines[1][i : i+4]
		s += lines[2][i : i+4]
		s += lines[3][i : i+4]
		fmt.Println(s)
		v += ParseNumber(s)
		v *= 10
	}
	return v / 10
}

func main() {
	var n int
	lines := make([]string, 5, 5)
	fmt.Scan(&n)
	fmt.Scan(&lines[0])
	fmt.Scan(&lines[1])
	fmt.Scan(&lines[2])
	fmt.Scan(&lines[3])
	fmt.Scan(&lines[4])

	fmt.Println(ReadNumber(lines, n))
}

/*
func ParseNumber(s []string) int {
	if s[0] == ".#.#" {
		return 4
	}
	if s[1] == ".##." {
		return 1
	}
	if s[3] == ".#.." {
		return 2
	}
	// s[0] == ".###"
	if s[1] == "...#" {
		if s[2] == "...#" {
			return 7
		} else {
			// s[3] == "...#"
			return 3
		}
	} else if s[1] == ".#.." {
		if s[3] == ".###" {
			return 5
		} else {
			// .#.#
			return 6
		}
	}
	return 0
}
*/
