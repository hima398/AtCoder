package main

import (
	"bufio"
	"errors"
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

type Item struct {
	W int
	V int
}

func Stupid(N, W int, A []int) int {

	return 0
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	W := nextInt()
	//fmt.Println(N, W)

	A := make([]Item, N+1)
	for i := 0; i < N; i++ {
		A[i].W = nextInt()
		A[i].V = nextInt()
	}
	// N番目の品物の中から重さWを超えないように選んだ価値の総和の最大
	DP := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		DP[i] = make([]int, W+1)
	}
	for j := A[0].W; j <= A[0].W+3; j++ {
		for i := 0; i < N; i++ {
			if A[i].W <= j {
				DP[i+1][j] = Max(DP[i][W-A[i].W]+A[i].V, DP[i][j])
			} else {
				DP[i+1][j] = DP[i][j]
			}
		}
	}
	fmt.Println(DP)
	fmt.Println(DP[N][A[0].W+3])

}
