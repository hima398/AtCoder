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

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

type Pos struct {
	X int
	Y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, K := nextInt(), nextInt()

	a := make([]int, N+1)
	s := make([]int, N+1)
	a[0] = 0
	s[0] = a[0]
	for i := 1; i <= N; i++ {
		a[i] = nextInt()
		s[i] = s[i-1] + a[i]
	}

	result := 0
	for i := 0; i < N-K+1; i++ {
		result += s[i+K] - s[i]
	}
	fmt.Println(result)

}
