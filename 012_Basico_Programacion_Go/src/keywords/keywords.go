package keywords

import "fmt"

func KeywordExample() {
	//defer
	// Se ejecuta de ultimas cuando todo el flujo halla sido recorrido
	defer fmt.Println("hola")
	fmt.Println("Mundo")

// Continue y break
	for i := 0; i < 10; i++{
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 4 {
			fmt.Println("Break")
			break
		}
}

}