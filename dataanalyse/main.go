package main

import (
	"fmt"
	"github.com/xiaoxiaoyijian123456/stats"
)

func main() {
	data := []float64{
		42, 70, 64, 47, 66, 55, 85, 10, 24, 45,
		16, 40, 81, 15, 35, 38, 79, 35, 36, 23,
		31, 38, 52, 16, 81, 69, 74, 38, 48, 25,
		31, 62, 47, 63, 84, 17, 40, 36, 44, 17,
		64, 75, 53, 31, 60, 12, 61, 43, 30, 33,
	}

	summary := stats.Summary(data)
	fmt.Printf("summary=%s\n", Json(summary))
}
