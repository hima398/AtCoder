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

func divide(n, x int) int {
	if n < 0 {
		return 0
	}
	return n/x + 1
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	a, b, x := nextInt(), nextInt(), nextInt()

	result := divide(b, x) - divide(a-1, x)
	fmt.Println(result)
}
