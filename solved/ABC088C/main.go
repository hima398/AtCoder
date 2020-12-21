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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	const n = 3
	var c [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = nextInt()
		}
	}
	var a [n]int
	var b [n]int
	a[0] = 0
	for i := 0; i < n; i++ {
		b[i] = c[0][i] - a[0]
	}
	for i := 1; i < n; i++ {
		a[i] = c[i][0] - b[0]
	}
	result := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result = result && (c[i][j] == a[i]+b[j])
			if !result {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
	return
}
