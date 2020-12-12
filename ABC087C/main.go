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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func FirstSubmission(n int, a1, a2 []int) int {
	result := 0
	for i := 0; i < n; i++ {
		total := 0
		for j := 0; j <= i; j++ {
			total += a1[j]
		}
		for k := i; k < n; k++ {
			total += a2[k]
		}
		if result < total {
			result = total
		}
	}
	return result
}

func DP(n int, a1, a2 []int) int {
	var dp [2][100]int

	dp[0][0] = a1[0]
	dp[1][0] = dp[0][0] + a2[0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + a1[i]
		dp[1][i] = Max(dp[0][i]+a2[i], dp[1][i-1]+a2[i])
	}
	return dp[1][n-1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	a1 := make([]int, n)
	a2 := make([]int, n)
	for i := 0; i < n; i++ {
		a1[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		a2[i] = nextInt()
	}

	//	fmt.Println(FirstSubmission(n, a1, a2))
	fmt.Println(DP(n, a1, a2))
}
