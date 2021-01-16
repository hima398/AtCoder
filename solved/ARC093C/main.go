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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n+2)
	s := 0
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		s += Abs(a[i-1] - a[i])
	}
	s += Abs(a[n] - a[n+1])
	for i := 1; i <= n; i++ {
		ans := s - Abs(a[i]-a[i+1]) - Abs(a[i-1]-a[i]) + Abs(a[i-1]-a[i+1])
		fmt.Println(ans)
	}
}
