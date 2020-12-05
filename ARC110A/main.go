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

type Pos struct {
	X int
	Y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	Max := 10000000000000
	var result int
	if N%2 == 0 {
		result = N + 1
	} else {
		result = N
	}
	for result <= Max {
		ok := true
		for j := 2; j <= N; j++ {
			ok = ok && result%j == 1
			if !ok {
				break
			}
		}
		if ok {
			break
		}
		result += 2
	}
	fmt.Println(result)
}
