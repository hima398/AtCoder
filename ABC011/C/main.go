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

func DP(N int, NG [301]bool) string {

	dp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = 2147483647
	}
	dp[N] = 0
	for i := N; i >= 0; i-- {
		if NG[i] {
			continue
		}
		for j := 1; j <= 3; j++ {
			if i-j < 0 {
				break
			}
			dp[i-j] = Min(dp[i]+1, dp[i-j])
		}
	}
	//fmt.Println(dp[0])
	if dp[0] <= 100 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()

	var NG [301]bool
	const NumNg = 3
	for i := 0; i < NumNg; i++ {
		ng := nextInt()
		if N == ng {
			fmt.Println("NO")
			return
		}
		NG[ng] = true
	}

	fmt.Println(DP(N, NG))
}
