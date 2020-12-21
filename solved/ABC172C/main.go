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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, M, K := nextInt(), nextInt(), nextInt()
	A := make([]int, N)
	B := make([]int, M)

	SA := make([]int, N+1)
	SB := make([]int, M+1)

	for i := 0; i < N; i++ {
		A[i] = nextInt()
		SA[i+1] = SA[i] + A[i]
		//		if SA[i] > K {
		//			break
		//		}
	}
	for i := 0; i < M; i++ {
		B[i] = nextInt()
		SB[i+1] = SB[i] + B[i]
		//		if SB[i] > K {
		//			break
		//		}
	}

	//if K > SA[N]+SB[M] {
	//	fmt.Println(N + M)
	//	return
	//}

	ans := 0
	j := M
	for i := 0; i <= N; i++ {
		if SA[i] > K {
			break
		}
		KA := K - SA[i]
		for SB[j] > KA {
			j--
		}
		ans = Max(ans, i+j)
	}

	fmt.Println(ans)
}
