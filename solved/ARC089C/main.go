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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	tn, xn, yn := 0, 0, 0
	n := nextInt()
	ok := true
	for i := 0; i < n; i++ {
		t, x, y := nextInt(), nextInt(), nextInt()
		ok = ok && (t%2 == (x+y)%2)
		ok = ok && (t-tn >= Abs(x-xn)+Abs(y-yn))
		if !ok {
			break
		}
		tn, xn, yn = t, x, y
	}
	if ok {
		fmt.Println("Yes")
		return
	} else {
		fmt.Println("No")
		return
	}

}
