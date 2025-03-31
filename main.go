package main
import "fmt"

func main()  {
	fmt.Println("Ucze sie")
	name := "Asia"
	age := 25
	height := 1.43
	isStudent := true
	weight := 1.68
	

	fmt.Println("Czesc", name)
	fmt.Println("Twoj wzrost", height)
	fmt.Println("Twoj wiek", age)
	fmt.Println("Czy student", isStudent)

	przywitaj(name)
	wynik := dodaj(2,4)
	fmt.Println("2+4= ", wynik)
	bmi_calc := bmi(weight, height)
	fmt.Printf("BMI: %.2f\n", bmi_calc)
}

func przywitaj(imie string) {
	fmt.Println("Czesc", imie,"milo sie widziec")
}

func dodaj(a int, b int) int {
	return a + b
}

func bmi (w float64, h float64) float64 { 
	return w / h * 2 * 2
}
