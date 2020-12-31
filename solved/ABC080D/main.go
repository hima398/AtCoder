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

func SolveCountChannel(N, C int) int {
	const debug = false
	var imos [31][100002]int
	channels := make([]int, 100002)
	max := 0
	for i := 0; i < N; i++ {
		S := nextInt()
		T := nextInt()
		CI := nextInt()
		imos[CI][S]++
		imos[CI][T+1]--
		max = Max(max, T)
	}
	for i := 1; i <= C; i++ {
		for j := 0; j < max; j++ {
			imos[i][j+1] += imos[i][j]
			if imos[i][j+1] > 0 {
				channels[j+1] |= 1 << i
			}
			if debug {
				fmt.Printf("%d, ", imos[i][j+1])
			}
		}
		if debug {
			fmt.Println("")
		}
	}
	if debug {
		fmt.Println("")
	}
	ans := 0
	for i := 1; i <= max; i++ {
		a := 0
		for j := 1; j <= C; j++ {
			a += channels[i] >> j & 1
		}
		ans = Max(ans, a)
		if debug {
			fmt.Printf("%d, ", channels[i])
		}
	}
	if debug {
		fmt.Println("")
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, C := nextInt(), nextInt()
	fmt.Println(SolveCountChannel(N, C))
	//fmt.Println(SolveOptimizeInput(N, C))
}
