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

func SolveImos(N, T int) int {
	Table := make([]int, 1100001)
	AN := 0
	for i := 0; i < N; i++ {
		A := nextInt()
		Table[A]++
		Table[A+T]--
		AN = Max(AN, A)
	}
	ans := 0
	for i := 0; i < AN+T+1; i++ {
		Table[i+1] += Table[i]
		if Table[i+1] > 0 {
			ans++
		}
	}
	return ans
}

func Solve(N, T int) int {
	ans := 0
	Prev := nextInt()
	for i := 1; i < N; i++ {
		A := nextInt()
		if Prev+T < A {
			// Prev < Prev + T < A
			ans += T
		} else {
			// Prev < A < Prev + T
			ans += A - Prev
		}
		Prev = A
	}
	ans += T
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, T := nextInt(), nextInt()
	fmt.Println(Solve(N, T))
}
