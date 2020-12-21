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

	N := nextInt()

	// Pilar
	P := make([]int, N)
	for i := 0; i < N; i++ {
		P[i] = nextInt()
	}

	dp := make([]int, N)
	dp[0] = 0
	dp[1] = dp[0] + Abs(P[1]-P[0])
	if N == 2 {
		fmt.Println(dp[1])
		return
	}
	for i := 2; i < N; i++ {
		dp[i] = Min(dp[i-1]+Abs(P[i]-P[i-1]), dp[i-2]+Abs(P[i]-P[i-2]))
	}
	fmt.Println(dp[N-1])
}
