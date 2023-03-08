package main

import "fmt"

func main() {
	// Goにおいては、オブジェクト指向などにおけるclassの概念がない
	// 似た役割として関連する情報をまとめるstruct(構造体)を使用する

	// 1.変数定義後にフィールドを設定する方法
	var user Person
	user.firstName = "testユーザーA"
	user.age = 40
	fmt.Println(user.firstName, user.age) 

	// 2.{} で順番にフィールドの値を渡す方法
	user2 := Person{"testユーザーB", 41}
	fmt.Println(user2.firstName, user2.age) 

	// 3.フィールド名を ： で指定する方法
	user3 := Person{age: 42, firstName:"testユーザーC"}
	fmt.Println(user3.firstName, user3.age) 

	// classの概念花が、構造体内に関数を定義できる
	// 通常の関数とこのなるのは、レシーバ引数の部分のみとなる

	/**
	 * func (<レシーバ引数>) <関数名>([引数]) [戻り値の型] {
	 * 	[関数の本体]
	 * }
	 */

	 // greetメソッドは p という名前の Person 型のレシーバを持つ
	 user4 := Person{age: 43, firstName:"testユーザーD"}
	 fmt.Println(user4.greet("Hello!")) // Hello! I am testユーザーD
}

type Person struct {
	firstName string
	age int
}


func (p Person) greet(message string) string {
	return message + " I am " + p.firstName
}