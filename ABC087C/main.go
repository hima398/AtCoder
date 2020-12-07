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
	fmt.Println(result)
}
