package main

import "fmt"
import "strconv"

func main() {

	// 条件分岐if 条件文を()で囲む必要がない
	var x int = 1

	if x == 1 {
		println("xは1")
	} else if x == 2 {
		println("xは2")
	} else {
		println("xは1,2以外")
	}

	// 条件分岐switch case文の後にbreakが不要
	// caseを跨ぐ場合、fallthroghを使用←要調査
	switch x {
		case 0:
			fmt.Println("xは0")
		case 1:
			fmt.Println("xは1")
		default:
			fmt.Println("xは1,2以外")
	}

	// caseに式が使える
	switch {
		case x == 1:
			fmt.Println("xは1")
	}

	// 繰り返しはforのみ
	for i := 0; i <= 10; i += 1 {
		// 参考 整数→文字列のキャストはstrconvのItoa関数を使用する(itoa とは「Integer to ASCII」の意)
		fmt.Println("iは" + strconv.Itoa(i))
	}


	var j int = 0
	// 継続条件のみも可能
	for j <= 10 {
		fmt.Println("jは" + strconv.Itoa(j))
		j++
	}
	
	// rangeを使用した繰り返しも可能
	/**
	var k int = 0
	for k, v:= range []int{0,1,2,3} {
		fmt.Println("kは" + strconv.Itoa(k))
		k++
	}
	*/
	

}