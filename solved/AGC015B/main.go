package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}

type Pos struct {
	X int
	Y int
}

func SolveStupidly(S string) int {
	ans := 0
	for i, _ := range S {
		for j, _ := range S {
			if i == j {
				continue
			}
			if i > j {
				if S[i] == 'U' {
					ans += 2
				} else {
					// S[i] == 'D'
					ans++
				}
			}
			if i < j {
				if S[i] == 'D' {
					ans += 2
				} else {
					ans++
				}
			}
		}
	}
	return ans
}

func Solve(S string) int {
	ans := 0
	for i := 1; i <= len(S); i++ {
		if S[i-1] == 'U' {
			ans += (len(S) - 1) + (i - 1)
		} else {
			//S[i-1] == 'D'
			ans += (len(S) - 1) + (len(S) - i)
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	S := nextString()
	fmt.Println(Solve(S))

}
