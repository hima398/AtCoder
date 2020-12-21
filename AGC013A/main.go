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

	n := nextInt()
	if n == 1 {
		fmt.Println(1)
		return
	}
	a := make([]int, n)
	a[0] = nextInt()
	a[1] = nextInt()
	var status int
	if a[1] > a[0] {
		status = 1
	} else if a[1] < a[0] {
		status = 2
	} else {
		status = 0
	}
	count := 1
	for i := 1; i < n-1; i++ {
		a[i+1] = nextInt()
		if status == 1 {
			if a[i+1] >= a[i] {
				continue
			} else {
				count++
				status = 0
			}
		} else if status == 2 {
			if a[i+1] <= a[i] {
				continue
			} else {
				count++
				status = 0
			}
		} else {
			if a[i+1] > a[i] {
				status = 1
				continue
			} else if a[i+1] < a[i] {
				status = 2
				continue
			} else {
				continue
			}
		}
	}
	fmt.Println(count)
}
