// run it with:
//    go run main.go
package main

import (
	"fmt"
)

func succ3n1(n int) (v int) {
	switch {
	case n == 1:
		v = 1
	case n%2 == 0:
		v = n / 2
	default:
		v = 3*n + 1
	}
	return
}

func cnt3n1(n int) (cnt int) {
	for {
		n = succ3n1(n)
		cnt += 1
		if n == 1 {
			break
		}
	}
	return
}

func list3n1(n int) (ret []int) {
	for {
		n := succ3n1(n)
		ret = append(ret, n)
		if n == 1 {
			break
		}
	}
	return
}

func maxlen3n1(m, n int) (k, maxlen int) {
	x := 0
	for i := m; i <= n; i++ {
		x = cnt3n1(i)
		if maxlen == 0 || x > maxlen {
			maxlen = x
			k = i
		}
	}
	return
}

func main() {
	i := 2
	j := 10
	k, maxlen := maxlen3n1(i, j)
	fmt.Printf("num = %d, maxlen=%d\n", k, maxlen)
}
