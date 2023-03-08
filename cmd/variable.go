package main

import "fmt"

func main() {

	// 変数宣言と代入
	var int int = 100
	fmt.Println(int)

	// 型を省略
	var n1 = 101
	fmt.Println(n1)

	
	// var n2 int エラーになる、なぜ？ → intだとエラー,int32だとエラーにならない
	// n2 = 102
	var n2 int32
	n2 = 102
	fmt.Println(n2)

	// varによる宣言を省略
	n3 := 103
	fmt.Println(n3)

	// varによる宣言をまとめる
	// n4 := 104と書くと、syntax error: unexpected :=, expected =
	var (
		n4 = 104
		n5 = 105
	) 
	fmt.Println(n4)
	fmt.Println(n5)

	// 使用していない変数はコンパイルエラー
	// m declared and not used
	// m := 0

	/*
	 * 整数: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune
	 * 浮動小数点: float32, float64
	 * 複素数: complex64, complex128
	 * 文字列: string
	 * 真偽値: bool
	 * インターフェース: error, any, comparable 
	 */

	var floatVar float32 = 106.001
	fmt.Println(floatVar)

	var boolVal bool = false
	fmt.Println(boolVal)

	// 変数はデフォルトで、数値は0、文字列は空文字、真偽値はfalse、インターフェースはnillになる

	// デフォルト値0
	var defaultInt int32
	fmt.Println(defaultInt)

	// デフォルト値""
	var defaultStr string
	fmt.Println(defaultStr)

	// デフォルト値false
	var defaultBool bool
	fmt.Println(defaultBool)

	// デフォルト値<nill>
	var defaultError error
	fmt.Println(defaultError)

	// 定数はconstで宣言
	const cst1 int32 = 107
	fmt.Println(cst1)

	// 型なしの定数も宣言可能
	const cst2 = 108
	fmt.Println(cst2)

	// 名前付き定数定義の右辺が省略できる
	// 省略された定数定義は、最後に定義した定数の右辺が設定される
	const (
		cst3 = 100 + 9
		cst4
		cst5
	)
	fmt.Println(cst3, cst4, cst5)

	// 演算子 +加算、-減産、/除算、*乗算、%余り
	
	// 代入演算子 =代入、:=変数の初期化と代入、+=、ｰ=演算と代入、++1加算、--1減算

	//　ビット演算子 |論理和、&論理積、^否定、<<左に算術シフト、>>右に算術シフト

	// 論理演算子 || または、 &&かつ、!否定

	// アドレス演算子 &ポインタを取得、*ポインタがさす値を取得(※ポインタ演算はできない)

	// チャネル演算 <- チャネルへの送受信


}