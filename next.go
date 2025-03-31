package main

import "fmt"

type Person struct { 
	Name string
	Height float64
	Age int
}

func main() { 
	asia := Person {
		Name: "Asia", 
		Height: 21.1,
		Age: 21,
	}
	fmt.Println(asia)
}
