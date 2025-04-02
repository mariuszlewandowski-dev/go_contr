package structgo
import fmt

func structgo() { 
	type Person struct {
		Name string
		Height float64
		Age int
	}
asia := Person{
	Name: "asia",
	Height: 132.2,
	Age: 43
}
fmt.Println(asia)
}