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

	N := nextInt()
	a := make([]int, N)
	m := make(map[int]int)

	for i := 0; i < N; i++ {
		a[i] = nextInt()
		m[a[i]-1]++
		m[a[i]]++
		m[a[i]+1]++
	}

	result := 0
	for i := 0; i < N; i++ {
		for j := i - 1; j <= i+1; j++ {
			count := m[a[j]]
			if result < count {
				result = count
			}
		}
	}
	fmt.Println(result)
}
