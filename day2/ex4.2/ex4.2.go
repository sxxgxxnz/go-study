package main

import "fmt"

func main() {

	var minimumWage int = 10 // 변수 minimumWage 선언 및 초기화
	var workingHour int = 20 // 변수 workingHour 선언 및 초기화

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)
}
