package main

import "fmt"

func main() {
	// 変数宣言
	var var1 float64 = 100

	// 変数のポインタ取得
	// &を使用して変数のアドレスを取得する
	// ポインタ型変数は*{変数の型}にて定義する
	var var1Pointer *float64 = &var1
	
	fmt.Println(var1) // 100
	fmt.Println(var1Pointer) // var1のポインタ

	// ポインタ型にて定義されている変数の値は*を付与することでアクセス可能
	fmt.Println(*var1Pointer) // var1のポインタの値

}