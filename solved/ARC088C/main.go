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

	x, y := nextInt(), nextInt()

	ans := 1
	a := x
	for {
		if y/2 >= a {
			a *= 2
			ans++
		} else {
			break
		}
	}
	fmt.Println(ans)
}
