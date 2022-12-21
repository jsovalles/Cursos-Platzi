package maps

import "fmt"

func MapsExample() {
	m := make(map[string]int)

	m["Jose"] = 12
	m["Pepe"] = 20

	fmt.Println(m)

	//Recorrer map
	for i,v := range m {
		fmt.Println(i,v)
	}

	//Encontrar valor
	value, ok := m["Jose"]
	println(value, ok)
}