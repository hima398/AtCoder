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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a := nextInt()
		m[a]++
	}
	//fmt.Println(m)
	ans := 0
	for k, v := range m {
		if v < k {
			ans += v
		} else {
			ans += v - k
		}
	}
	fmt.Println(ans)
}
