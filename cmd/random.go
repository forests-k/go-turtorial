package main

import "fmt"
import "math/rand"
import "time"

func main() {
	/**
	 コンピュータ上では完全な乱数は作れない
	 乱数っぽい並びの数列を作る
	 Goではmath/randパッケージを使う
 	 乱数の生成には元（種）になる数値が必要
	 種は現在時刻を使うことが多い（毎回変わるので）→timeパッケージを使う
	 */

	// 現在時刻を数値で取得
	t := time.Now().UnixNano()
	// 乱数の種を設定
	rand.Seed(t)
	// 0~9の乱数を発生
	s := rand.Intn(10)
	fmt.Println(s)
}