package main

import (
	"fmt"
	"lesson_go/tools/bmi"
)

func main() {
	result := bmi.Calculate(64, 20)
	fmt.Println(result)
	test0 := bmi.Hello("onetwothree")
	fmt.Println(test0)

}
