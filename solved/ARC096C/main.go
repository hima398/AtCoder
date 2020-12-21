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

func FullSearch(a, b, c, x, y int) int {
	const N = 100000
	result := 2147483647
	for i := 0; i <= N; i++ {
		total := i*2*c + Max(0, x-i)*a + Max(0, y-i)*b
		if total < result {
			result = total
		}
		if i > x+y {
			break
		}
	}
	return result
}

func O1(a, b, c, x, y int) int {
	//ハーフを買わない
	if a+b < 2*c {
		return a*x + b*y
	}

	// ハーフのみで買う
	halfOnly := c * Max(x, y) * 2
	// 1種類ハーフ、A or Bの残りを買い足す
	var mix int
	if x > y {
		mix = 2*c*y + a*(x-y)
	} else {
		mix = 2*c*x + b*(y-x)
	}

	return Min(halfOnly, mix)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c, x, y := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	//fmt.Println(FullSearch(a, b, c, x, y))
	fmt.Println(O1(a, b, c, x, y))
}
