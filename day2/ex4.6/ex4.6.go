package main

import "fmt"

var g int = 10 //패키지 전역 변수 선언

func main() {

	var m int = 20 //로컬 변수 선언

	{
		var s int = 50 //로컬 변수 선언
		fmt.Println(m, s, g)
	} //s 로컬변수는 사라짐

	//m = s + 20 //error
}
