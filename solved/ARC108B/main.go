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
	s := nextString()

	if n < 3 {
		fmt.Println(len(s))
		return
	}
	t := s[:2]
	s = s[2:]
	for {
		t = t + string(s[0])
		s = s[1:]
		//fmt.Println(t, s)
		if len(t) >= 3 && t[len(t)-3:] == "fox" {
			t = t[:len(t)-3]
		}
		if s == "" {
			break
		}
	}
	//fmt.Println(t, s)
	fmt.Println(len(t))
}
