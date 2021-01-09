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

	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	if a == b+c+d || b == a+c+d || c == a+b+d || d == a+b+c {
		fmt.Println("Yes")
		return
	}
	if a+b == c+d || a+c == b+d || a+d == b+c {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
