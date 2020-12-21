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

func Max(a, b, c int) int {
	if a < b {
		if b < c {
			// a < b < c
			return c
		} else {
			// c < b
			return b
		}
	} else {
		if a < c {
			// b =< a < c
			return c
		} else {
			// c < a
			return a
		}
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c := nextInt(), nextInt(), nextInt()
	m := Max(a, b, c)
	total := a + b + c
	if (3*m-total)%2 == 0 {
		fmt.Println((3*m - total) / 2)
	} else {
		fmt.Println((3*(m+1) - total) / 2)
	}
}
