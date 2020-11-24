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

func TrainTicket(a, b, c, d int) string {
	if a-b-c-d == 7 {
		return fmt.Sprintf("%d-%d-%d-%d=7", a, b, c, d)
	} else if a-b-c+d == 7 {
		return fmt.Sprintf("%d-%d-%d+%d=7", a, b, c, d)
	} else if a-b+c-d == 7 {
		return fmt.Sprintf("%d-%d+%d-%d=7", a, b, c, d)
	} else if a-b+c+d == 7 {
		return fmt.Sprintf("%d-%d+%d+%d=7", a, b, c, d)
	} else if a+b-c-d == 7 {
		return fmt.Sprintf("%d+%d-%d-%d=7", a, b, c, d)
	} else if a+b-c+d == 7 {
		return fmt.Sprintf("%d+%d-%d+%d=7", a, b, c, d)
	} else if a+b+c-d == 7 {
		return fmt.Sprintf("%d+%d+%d-%d=7", a, b, c, d)
	} else {
		// a + b + c + d == 7
		return fmt.Sprintf("%d+%d+%d+%d=7", a, b, c, d)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	abcd := nextInt()
	a := abcd / 1000
	b := (abcd - 1000*a) / 100
	c := (abcd - (1000*a + 100*b)) / 10
	d := abcd - (1000*a + 100*b + 10*c)
	fmt.Println(TrainTicket(a, b, c, d))
}
