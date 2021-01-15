package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Pow(x, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}
	return ret
}

func Solve(n int) (int, int, bool) {
	for a := 1; a <= 37; a++ {
		for b := 1; b <= 25; b++ {
			if n-Pow(3, a) == Pow(5, b) {
				return a, b, true
			}
		}
	}
	return -1, -1, false
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a, b, ok := Solve(n)
	if ok {
		fmt.Println(a, b)
	} else {
		fmt.Println(-1)
	}
}
