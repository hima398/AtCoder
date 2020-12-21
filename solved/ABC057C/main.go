package main

import (
	"bufio"
	"fmt"
	"math"
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
	if x > y {
		return x
	}
	return y
}

func F(a, b int) int {
	return Max(len(strconv.Itoa(a)), len(strconv.Itoa(b)))
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	sqrtN := int(math.Sqrt(float64(n)))

	// N <= 10 ^ 10
	result := 11
	for a := 1; a < sqrtN+1; a++ {
		if n%a != 0 {
			continue
		}
		f := F(a, n/a)
		if f < result {
			result = f
		}
	}
	fmt.Println(result)
}
