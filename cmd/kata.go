package main

import "fmt"

func main() {
	// 単純な数値でのキャスト
	var f float64 = 100
	var n int = int(f)
	fmt.Println(n)

	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		println("good")
	}

}