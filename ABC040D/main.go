package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
	size []int
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.par[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (this *UnionFind) Find(x int) int {
	if this.par[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		this.par[x] = this.Find(this.par[x])
		return this.par[x]
	}
}

func (this *UnionFind) ExistSameUnion(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Unite(x, y int) {
	rootX := this.Find(x)
	rootY := this.Find(y)
	if rootX == rootY {
		return
	}

	if this.size[rootX] < this.size[rootY] {
		this.par[rootX] = rootY
		this.size[rootY] += this.size[rootX]
	} else {
		//this.size[rootX] >= this.size[rootY]
		this.par[rootY] = rootX
		this.size[rootX] += this.size[rootY]
	}
	// raink
	/*
		if this.rank[x] < this.rank[y] {
			this.par[x] = y
		} else {
			// this.rank[x] >= this.rank[y]
			this.par[y] = x
			if this.rank[x] == this.rank[y] {
				this.rank[x]++
			}
		}
	*/
}

func (this *UnionFind) Size(x int) int {
	return this.size[this.Find(x)]
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.par)
	fmt.Println(u.rank)
	fmt.Println(u.size)
}

type Query struct {
	i    int
	mode int
	a    int
	b    int
	y    int
}

type Answer struct {
	i int
	v int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	u := NewUnionFind(n)
	var queries []*Query
	for i := 0; i < m; i++ {
		a, b, y := nextInt(), nextInt(), nextInt()
		queries = append(queries, &Query{i, 0, a, b, y})
	}
	q := nextInt()
	for i := 0; i < q; i++ {
		a, y := nextInt(), nextInt()
		queries = append(queries, &Query{i, 1, a, -1, y})
	}
	sort.Slice(queries, func(i, j int) bool {
		if queries[i].y > queries[j].y {
			return true
		}
		if queries[i].y < queries[j].y {
			return false
		}
		// queries[i].y == queries[j].y
		return queries[i].mode > queries[j].mode
	})
	var answers []Answer
	for _, v := range queries {
		//fmt.Println(v)
		if v.mode == 0 {
			u.Unite(v.a, v.b)
		} else if v.mode == 1 {
			ans := Answer{v.i, u.Size(v.a)}
			answers = append(answers, ans)
		}
	}
	sort.Slice(answers, func(i, j int) bool {
		return answers[i].i < answers[j].i
	})
	for _, ans := range answers {
		fmt.Println(ans.v)
	}
}
