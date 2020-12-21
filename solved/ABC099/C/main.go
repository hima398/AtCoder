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

var memo map[int]int
/*
func GreedyAlgorithm(N int) int {
	a := []int{1, 6, 9, 36, 81, 216, 729, 1296, 6561, 7776, 46656, 59049}

	b := len(a) - 1
	for a[b] > N {
		b--
	}

	result := 0
	for N > 0 {
		if N < a[b] {
			b--
			continue
		}
		fmt.Println(N, a[b])
		N -= a[b]
		result++
	}
	return result
}
*/

func DP(N int) int {
	a := []int{1, 6, 9, 36, 81, 216, 729, 1296, 6561, 7776, 46656, 59049}
	var dp [100001]int

	dp[0] = 0
	for i := 1; i < N+1; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < len(a); j++ {
			if i-a[j] < 0 {
				break
			}
			dp[i] = Min(dp[i], dp[i-a[j]]+1)
		}
	}
	return dp[N]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()

	fmt.Println(DP(N))
}
