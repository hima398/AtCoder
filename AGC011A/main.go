package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Pos struct {
	X int
	Y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, c, k := nextInt(), nextInt(), nextInt()
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = nextInt()
	}
	sort.Ints(t)

	cap := 1
	bus := 1
	first := t[0]
	for i := 0; i < n; i++ {
		if cap > c || t[i] > first+k {
			bus++
			cap = 1
			first = t[i]
		}
		cap++
	}
	fmt.Println(bus)
}
