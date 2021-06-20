package main

import "fmt"

func main() {
	var a int = 10 //var '변수' '타입' 형태로 선언한다.
	var msg string = "hello variable"

	//변수 값 변경
	a = 20
	msg = "good"

	fmt.Println(msg, a)

}
