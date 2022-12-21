package structs

import "fmt"

type car struct {
	brand string
	year  int
}

func StructsExample() {

	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}