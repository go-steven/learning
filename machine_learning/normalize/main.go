package main

import (
	"fmt"
	"github.com/cdipaolo/goml/base"
)

func main() {
	u := []float64{1.0, 2.0, 3.0, 4.2, 5.7, 8.3}
	v := []float64{1.2, 2.0, 3.3, 4.0, 5.2, 8.1}
	fmt.Printf("u=%v\n", u)
	fmt.Printf("v=%v\n", v)

	fmt.Printf("EuclideanDistance = %v\n", base.EuclideanDistance(u, v)) // 欧氏距离
	fmt.Printf("ManhattanDistance = %v\n", base.ManhattanDistance(u, v)) // 曼哈顿距离

	fmt.Printf("LNorm(1) = %v\n", base.LNorm(1)(u, v)) // 皮尔逊系数, p=1
	fmt.Printf("LNorm(2) = %v\n", base.LNorm(2)(u, v)) // 皮尔逊系数, p=2
	fmt.Printf("LNorm(3) = %v\n", base.LNorm(3)(u, v)) // 皮尔逊系数, p=3
	fmt.Printf("LNorm(4) = %v\n", base.LNorm(4)(u, v)) // 皮尔逊系数, p=4
	fmt.Printf("LNorm(5) = %v\n", base.LNorm(5)(u, v)) // 皮尔逊系数, p=5
	fmt.Printf("LNorm(6) = %v\n", base.LNorm(6)(u, v)) // 皮尔逊系数, p=6

	base.NormalizePoint(u)
	base.NormalizePoint(v)
	fmt.Printf("u=%v\n", u)
	fmt.Printf("v=%v\n", v)
	
	fmt.Printf("EuclideanDistance = %v\n", base.EuclideanDistance(u, v)) // 欧氏距离
	fmt.Printf("ManhattanDistance = %v\n", base.ManhattanDistance(u, v)) // 曼哈顿距离

	fmt.Printf("LNorm(1) = %v\n", base.LNorm(1)(u, v)) // 皮尔逊系数, p=1
	fmt.Printf("LNorm(2) = %v\n", base.LNorm(2)(u, v)) // 皮尔逊系数, p=2
	fmt.Printf("LNorm(3) = %v\n", base.LNorm(3)(u, v)) // 皮尔逊系数, p=3
	fmt.Printf("LNorm(4) = %v\n", base.LNorm(4)(u, v)) // 皮尔逊系数, p=4
	fmt.Printf("LNorm(5) = %v\n", base.LNorm(5)(u, v)) // 皮尔逊系数, p=5
	fmt.Printf("LNorm(6) = %v\n", base.LNorm(6)(u, v)) // 皮尔逊系数, p=6
}
