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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	m := make(map[int]bool)
	i := 1
	var result int
	for {
		result++
		a := a[i-1]
		if m[a] {
			fmt.Println(-1)
			return
		}
		if a == 2 {
			break
		}
		m[a] = true
		i = a
	}
	fmt.Println(result)
}
